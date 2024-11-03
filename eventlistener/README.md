# Event Listener

This package is a tool for listening to Smart Contract events.

## How to Use

Please refer to the example in [examples/eventlistener](../examples/eventlistener).

## Notable Logic

1. The event listener can accept multiple websocket URLs and will listen to all of them concurrently
2. The event listener will do best-effort deduplication of events using in-memory cache
   to prevent processing the same event multiple times
3. The event listener will start listening with at least 1 of the websocket URL working
   (it doesn't require all provided websocket URLs to work)

## Options

There are a few options available on the `eventlistener.New` function.

- `WithMaxConcurrentProcess`: Set the maximum number of concurrent process. Default is 100.
