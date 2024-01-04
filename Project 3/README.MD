## NATS Request/Reply simple example

### Setup:

```
go mod init nats-example
go get github.com/nats-io/nats.go/
```

### Start:

```
go run responder.go
go run requester.go
```
