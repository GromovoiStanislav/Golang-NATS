## NATS Basic Usage simple example

### Setup:

```
go mod init nats-example
go get github.com/nats-io/nats.go/
```

### Start:

```
go run example-1.go // Async Subscriber
go run example-2.go // Sync Subscriber
go run example-3.go // Channel Subscriber
go run example-4.go // Requests/Replies
```
