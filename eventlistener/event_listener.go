package eventlistener

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/patrickmn/go-cache"
	"github.com/pixel8labs/logtrace/log"
	"github.com/pixel8labs/logtrace/trace"
)

type HandlerFn func(ctx context.Context, msg types.Log)

// EventListener listens for events from the smart contract.
// This is already equipped with log & trace using github.com/pixel8labs/logtrace package.
type EventListener struct {
	appName              string
	maxConcurrentProcess int
	websocketUrls        []string
	contractAddress      common.Address
	handlers             map[common.Hash]HandlerFn
	deduplicationCache   *cache.Cache
	cancelFunc           context.CancelFunc
}

type NewOption func(*EventListener)

// WithMaxConcurrentProcess sets the maximum number of concurrent processes for each websocket URL. The default is 100.
func WithMaxConcurrentProcess(max int) NewOption {
	return func(e *EventListener) {
		e.maxConcurrentProcess = max
	}
}

// New creates a new EventListener.
// websocketUrls is a list of websocket URLs to connect to the Ethereum node.
// The expectation is at least one URL is provided. If multiple URLs are provided,
// the event listener will try to connect & listen to all of them, and expect at least one of them is working.
// The event listener will also perform best-effort deduplication using in-memory cache.
func New(
	appName string,
	websocketUrls []string,
	contractAddress common.Address,
	opts ...NewOption,
) *EventListener {
	e := &EventListener{
		appName:              appName,
		maxConcurrentProcess: 100,
		websocketUrls:        websocketUrls,
		contractAddress:      contractAddress,
		handlers:             make(map[common.Hash]HandlerFn),
		// The deduplication cache only needs to be very short-lived.
		deduplicationCache: cache.New(1*time.Minute, 2*time.Minute),
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// RegisterHandler registers a new event handler.
func (e *EventListener) RegisterHandler(eventHash common.Hash, fn HandlerFn) {
	e.handlers[eventHash] = fn
}

// Listen starts listening for events. To gracefully stop, call Stop().
func (e *EventListener) Listen(ctx context.Context) error {
	// Connect to all the Ethereum nodes.
	var mapWsUrlToEthClient = make(map[string]*ethclient.Client)
	for _, url := range e.websocketUrls {
		ethClient, err := ethclient.Dial(url)
		if err != nil {
			log.Error(context.Background(), err, log.Fields{
				"url": url,
			}, "EventListener: failed to connect to Ethereum node")
			continue
		}
		mapWsUrlToEthClient[url] = ethClient
	}

	// If no Ethereum node is connected, return error.
	if len(mapWsUrlToEthClient) == 0 {
		return fmt.Errorf("EventListener: no Ethereum node is connected")
	}

	ctx, e.cancelFunc = context.WithCancel(ctx)

	// Start listening for events.
	var wg sync.WaitGroup
	var errs []error
	for wsUrl, ethClient := range mapWsUrlToEthClient {
		wg.Add(1)
		go func() {
			if err := e.runListener(ctx, wsUrl, ethClient); err != nil {
				log.Error(ctx, err, log.Fields{
					"url": wsUrl,
				}, "EventListener: failed to run listener")
				errs = append(errs, err)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// If all listener returns error, return all errors.
	if len(errs) == len(mapWsUrlToEthClient) {
		return fmt.Errorf("all listeners failed to start listening: %v", errs)
	}

	return nil
}

func (e *EventListener) runListener(ctx context.Context, websocketUrl string, ethClient *ethclient.Client) error {
	logs := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{
			e.contractAddress,
		},
	}, logs)
	if err != nil {
		return fmt.Errorf("ethClient.SubscribeFilterLogs: %w", err)
	}
	defer sub.Unsubscribe()

	logFields := log.Fields{
		"contract_address":       e.contractAddress.String(),
		"max_concurrent_process": e.maxConcurrentProcess,
		"url":                    websocketUrl,
	}
	log.Info(ctx, logFields, "EventListener: listening for events...")

	// WaitGroup to wait until all process is done.
	var wg sync.WaitGroup
	// maxProcessCh is used to limit the number of concurrent process.
	maxProcessCh := make(chan int, e.maxConcurrentProcess)
	for {
		select {
		case err := <-sub.Err():
			if err != nil {
				return fmt.Errorf("subscription error: %w", err)
			}
		case msg := <-logs:
			wg.Add(1)
			maxProcessCh <- 1
			// Do process async.
			go func() {
				e.processLog(ctx, msg)
				<-maxProcessCh
				wg.Done()
			}()
		case <-ctx.Done():
			// Wait until all process is done.
			log.Info(ctx, logFields, "EventListener: received stop signal. Waiting for all processes to finish...")
			wg.Wait()
			// Close connection.
			ethClient.Close()
			log.Info(ctx, logFields, "EventListener: stopped")
			return nil
		}
	}
}

// Stop stops the listener.
func (e *EventListener) Stop() {
	e.cancelFunc()
}

func (e *EventListener) processLog(ctx context.Context, msg types.Log) {
	fn, ok := e.handlers[msg.Topics[0]]
	if !ok {
		// If no handler, just ignore and return.
		return
	}

	ctx, span := trace.StartSpan(
		ctx,
		e.appName+"-listener",
		msg.Topics[0].String(),
	)
	defer span.End()

	logFields := log.Fields{
		"event": msg,
	}

	// Deduplication logic. If the event is already processed, skip.
	if _, found := e.deduplicationCache.Get(string(msg.Data)); found {
		log.Info(ctx, logFields, "EventListener: duplicated event, skipping...")
		return
	} else {
		// Set the cache to true at the very beginning of the process to prevent
		// multiple incoming messages at the same time to be processed multiple times.
		e.deduplicationCache.Set(string(msg.Data), true, cache.DefaultExpiration)
	}

	log.Info(ctx, logFields, "EventListener: processing event...")
	fn(ctx, msg)
	log.Info(ctx, logFields, "EventListener: processed event")
}
