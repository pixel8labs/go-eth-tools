package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/pixel8labs/go-eth-tools/eventlistener"
	"github.com/pixel8labs/logtrace/log"
	"github.com/pixel8labs/logtrace/trace"
)

const appName = "erc20-event-listener"
const appEnv = "local"

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Init logger & tracer.
	log.Init(appName, appEnv)
	trace.InitTracer()

	websocketUrl := os.Getenv("WEBSOCKET_URL")
	contractAddress := common.HexToAddress(os.Getenv("ERC20_CONTRACT_ADDRESS"))

	// We can use either ABI or the generated Go code from the contract to unpack the event.
	erc20AbiJson, err := abi.JSON(strings.NewReader(erc20Abi))
	if err != nil {
		panic(err)
	}

	client, err := ethclient.Dial(websocketUrl)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Create a new event listener.
	eventListener := eventlistener.New(appName, client, contractAddress)

	// Register the handler function.
	eventListener.RegisterHandler(erc20AbiJson.Events["Transfer"].ID, func(ctx context.Context, msg types.Log) {
		content := make(map[string]any)
		if err := erc20AbiJson.UnpackIntoMap(content, "Transfer", msg.Data); err != nil {
			log.Error(ctx, err, log.Fields{
				"msg": msg,
			}, "Listener.Transfer: Failed to unpack event")
			return
		}

		log.Info(ctx, log.Fields{"content": content}, "Listener.Transfer: Event processed")
	})

	// We can register other handlers here for other events.

	// Do graceful shutdown.
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	done := make(chan struct{})

	go func() {
		if err := eventListener.Listen(ctx); err != nil {
			panic(err)
		}
		done <- struct{}{}
	}()

	<-ctx.Done()
	eventListener.Stop()
	<-done
	cancel()
}
