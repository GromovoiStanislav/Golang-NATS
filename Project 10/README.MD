## NATS JetStream simple example

### Setup:

```
go mod init nats-example
go get github.com/nats-io/nats.go/
go get github.com/nats-io/nats.go/jetstream@v1.31.0
```

### Start:

```
go run сonsumer.go
go run publisher.go
```
