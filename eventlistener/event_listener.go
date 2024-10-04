package eventlistener

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pixel8labs/logtrace/log"
	"github.com/pixel8labs/logtrace/trace"
)

type HandlerFn func(ctx context.Context, msg types.Log)

// EventListener listens for events from the smart contract.
// This is already equipped with log & trace using github.com/pixel8labs/logtrace package.
type EventListener struct {
	appName              string
	maxConcurrentProcess int
	ethClient            *ethclient.Client
	contractAddress      common.Address
	handlers             map[common.Hash]HandlerFn
	stopCh               chan struct{}
}

type NewOption func(*EventListener)

// WithMaxConcurrentProcess sets the maximum number of concurrent processes. The default is 100.
func WithMaxConcurrentProcess(max int) NewOption {
	return func(e *EventListener) {
		e.maxConcurrentProcess = max
	}
}

// WithAbi sets the ABI of the contract.
// Abi will be used to automatically registers & unpack the event data.
func WithAbi(abi string) NewOption {
	return func(e *EventListener) {
		// TODO
	}
}

// New creates a new EventListener.
func New(
	appName string,
	ethClient *ethclient.Client,
	contractAddress common.Address,
	opts ...NewOption,
) *EventListener {
	e := &EventListener{
		appName:              appName,
		maxConcurrentProcess: 100,
		ethClient:            ethClient,
		contractAddress:      contractAddress,
		stopCh:               make(chan struct{}),
		handlers:             make(map[common.Hash]HandlerFn),
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

// Listen starts listening for events.
func (e *EventListener) Listen(ctx context.Context) error {
	logs := make(chan types.Log)
	sub, err := e.ethClient.SubscribeFilterLogs(ctx, ethereum.FilterQuery{
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
			// Do process async
			go e.processLog(ctx, msg)
			<-maxProcessCh
			wg.Done()
		case <-e.stopCh:
			// Wait until all process is done.
			log.Info(ctx, logFields, "EventListener: received stop signal. Waiting for all processes to finish...")
			wg.Wait()
			close(e.stopCh)
			log.Info(ctx, logFields, "EventListener: stopped")
			return nil
		}
	}
}

func (e *EventListener) Stop() {
	e.stopCh <- struct{}{}
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
	log.Info(ctx, logFields, "EventListener: processing event...")
	fn(ctx, msg)
	log.Info(ctx, logFields, "EventListener: processed event")
}
