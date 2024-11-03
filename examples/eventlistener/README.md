# Event Listener Example

This is an example of how to use the `eventlistener` package.
We listen to ERC20 Transfer event as an example.

## How to Run

1. Copy `.env.example` to `.env` and fill in all the values.
2. Run the following command:

```bash
go run .
```

## Expected Output

```json
{"level":"info","context":{"contract_address":"0xbDa130737BDd9618301681329bF2e46A016ff9Ad","max_concurrent_process":100,"url":"wss://<url1>"},"service":"erc20-event-listener","env":"local","time":"2024-11-03T20:08:22+07:00","message":"EventListener: listening for events..."}
{"level":"info","context":{"contract_address":"0xbDa130737BDd9618301681329bF2e46A016ff9Ad","max_concurrent_process":100,"url":"wss://<url2>"},"service":"erc20-event-listener","env":"local","time":"2024-11-03T20:08:22+07:00","message":"EventListener: listening for events..."}
{"level":"info","context":{"event":{"address":"0xbda130737bdd9618301681329bf2e46a016ff9ad","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000002c1f148ee973a4cda4abece2241df3d3337b7319"],"data":"0x00000000000000000000000000000000000000000000000043ca593b0fc17348","blockNumber":"0x619aa9","transactionHash":"0xde36571e1ea78bcd56ee36acabb750438dbd661356b8e0f0356a308613d53c30","transactionIndex":"0x1","blockHash":"0x63654842e6c1ea2fae11b33eaeefc58a2998221ff76c3b9a06a383b20532dc99","logIndex":"0x4","removed":false}},"service":"erc20-event-listener","env":"local","trace_id":"2211fde018e9188587951d11b3cc4931","span_id":"3a57cf306b944272","time":"2024-11-03T20:08:24+07:00","message":"EventListener: duplicated event, skipping..."}
{"level":"info","context":{"event":{"address":"0xbda130737bdd9618301681329bf2e46a016ff9ad","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000007f1d8cd8120722f937338208412318d9c5d35e1f"],"data":"0x000000000000000000000000000000000000000000000000117219b6c32ad02c","blockNumber":"0x619aa9","transactionHash":"0xde36571e1ea78bcd56ee36acabb750438dbd661356b8e0f0356a308613d53c30","transactionIndex":"0x1","blockHash":"0x63654842e6c1ea2fae11b33eaeefc58a2998221ff76c3b9a06a383b20532dc99","logIndex":"0x3","removed":false}},"service":"erc20-event-listener","env":"local","trace_id":"6feb5c66b59f5b04e5443e6d5ba3b9ae","span_id":"94ae1a6445bbb843","time":"2024-11-03T20:08:24+07:00","message":"EventListener: processing event..."}
{"level":"info","context":{"content":{"value":1257095518738894892}},"service":"erc20-event-listener","env":"local","trace_id":"6feb5c66b59f5b04e5443e6d5ba3b9ae","span_id":"94ae1a6445bbb843","time":"2024-11-03T20:08:24+07:00","message":"Listener.Transfer: Event processed"}
{"level":"info","context":{"event":{"address":"0xbda130737bdd9618301681329bf2e46a016ff9ad","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000007f1d8cd8120722f937338208412318d9c5d35e1f"],"data":"0x000000000000000000000000000000000000000000000000117219b6c32ad02c","blockNumber":"0x619aa9","transactionHash":"0xde36571e1ea78bcd56ee36acabb750438dbd661356b8e0f0356a308613d53c30","transactionIndex":"0x1","blockHash":"0x63654842e6c1ea2fae11b33eaeefc58a2998221ff76c3b9a06a383b20532dc99","logIndex":"0x3","removed":false}},"service":"erc20-event-listener","env":"local","trace_id":"6feb5c66b59f5b04e5443e6d5ba3b9ae","span_id":"94ae1a6445bbb843","time":"2024-11-03T20:08:24+07:00","message":"EventListener: processed event"}
{"level":"info","context":{"contract_address":"0xbDa130737BDd9618301681329bF2e46A016ff9Ad","max_concurrent_process":100,"url":"wss://<url1>"},"service":"erc20-event-listener","env":"local","time":"2024-11-03T20:08:29+07:00","message":"EventListener: received stop signal. Waiting for all processes to finish..."}
{"level":"info","context":{"contract_address":"0xbDa130737BDd9618301681329bF2e46A016ff9Ad","max_concurrent_process":100,"url":"wss://<url2>"},"service":"erc20-event-listener","env":"local","time":"2024-11-03T20:08:29+07:00","message":"EventListener: received stop signal. Waiting for all processes to finish..."}
{"level":"info","context":{"contract_address":"0xbDa130737BDd9618301681329bF2e46A016ff9Ad","max_concurrent_process":100,"url":"wss://<url1>"},"service":"erc20-event-listener","env":"local","time":"2024-11-03T20:08:29+07:00","message":"EventListener: stopped"}
{"level":"info","context":{"contract_address":"0xbDa130737BDd9618301681329bF2e46A016ff9Ad","max_concurrent_process":100,"url":"wss://<url2>"},"service":"erc20-event-listener","env":"local","time":"2024-11-03T20:08:29+07:00","message":"EventListener: stopped"}
```
