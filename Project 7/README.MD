## NATS Request/Reply with Protobuf simple example

NATS message payloads are byte slices, so any kind of serialization strategy can be applied. This example shows a simple way to define message types using the Protocol Buffers IDL, generate code for the target language, and then use it with NATS.

#### Generate code:

```
protoc --go_out=. types.proto
```

#### Setup:

```
go mod init nats-example
go get github.com/nats-io/nats.go/
go get google.golang.org/protobuf/proto
```

#### Start:

```
go run .
```
