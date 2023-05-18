# natstern

> like stern but for NATS, except different

nats debug tool to list all the consumers for a stream and their pending messages. Gets your credentials from nats cli context.

## Installation

(you need Go installed)

```bash
go install github.com/robindiddams/natstern
```

## Usage

```bash
> natstern stream-name
using context production
consumer-number1 pending: 23, ackPending: 1
consumer-number2 pending: 27857, ackPending: 0
webhooks pending: 253, ackPending: 0

42 consumers found for stream 'stream-name'
consumers with zero pending messages omitted, use -z to view them
```
