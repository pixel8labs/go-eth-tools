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
{"level":"info","context":{"contract_address":"0xbDa130737BDd9618301681329bF2e46A016ff9Ad","max_concurrent_process":100},"service":"erc20-event-listener","env":"local","time":"2024-10-04T22:33:51+07:00","message":"EventListener: listening for events..."}
{"level":"info","context":{"event":{"address":"0xbda130737bdd9618301681329bf2e46a016ff9ad","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000002c1f148ee973a4cda4abece2241df3d3337b7319"],"data":"0x0000000000000000000000000000000000000000000000007b92f32f3cddc180","blockNumber":"0x4e928c","transactionHash":"0xf879515abf78aed522018490f7b246dcc96b3b59174f9ce32210cd572e4a06ac","transactionIndex":"0x0","blockHash":"0x4d89f3f8d10c19a5f498758e19d0a21f8be30333740ff404df4a304bf281a488","logIndex":"0x2","removed":false},"event_name":"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"},"service":"erc20-event-listener","env":"local","trace_id":"cd553121901b724b6d0d755cfa27f9d5","span_id":"cd4a5f9d74f0111d","time":"2024-10-04T22:33:51+07:00","message":"EventListener: processing event..."}
{"level":"info","context":{"event":{"address":"0xbda130737bdd9618301681329bf2e46a016ff9ad","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000002c1f148ee973a4cda4abece2241df3d3337b7319","0x000000000000000000000000e4303f01d1557ab99236f8678e2e522b0354f109"],"data":"0x0000000000000000000000000000000000000000000000000001f20fc7c23cc7","blockNumber":"0x4e928c","transactionHash":"0x9933d1c65197920778ed662e99e62b892b606308141646fe59cc17159319b8cb","transactionIndex":"0x27","blockHash":"0x4d89f3f8d10c19a5f498758e19d0a21f8be30333740ff404df4a304bf281a488","logIndex":"0x65","removed":false},"event_name":"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"},"service":"erc20-event-listener","env":"local","trace_id":"5d9de8661b93454063cf5da13da085dc","span_id":"82c31714cb2ed5e9","time":"2024-10-04T22:33:51+07:00","message":"EventListener: processing event..."}
{"level":"info","context":{"contract_address":"0xbDa130737BDd9618301681329bF2e46A016ff9Ad","max_concurrent_process":100},"service":"erc20-event-listener","env":"local","time":"2024-10-04T22:33:55+07:00","message":"EventListener: received stop signal. Waiting for all process to finish..."}
{"level":"info","context":{"contract_address":"0xbDa130737BDd9618301681329bF2e46A016ff9Ad","max_concurrent_process":100},"service":"erc20-event-listener","env":"local","time":"2024-10-04T22:33:55+07:00","message":"EventListener: stopped"}
```
