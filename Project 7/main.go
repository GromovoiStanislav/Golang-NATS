package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)


func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)
	defer nc.Drain()

	nc.Subscribe("greet", func(msg *nats.Msg) {
		var req GreetRequest
		proto.Unmarshal(msg.Data, &req)


		rep := GreetReply{
			Text: fmt.Sprintf("hello %q!", req.Name),
		}
		data, _ := proto.Marshal(&rep)
		msg.Respond(data)
	})

	req := GreetRequest{
		Name: "joe",
	}
	data, _ := proto.Marshal(&req)

	msg, _ := nc.Request("greet", data, time.Second)

	var rep GreetReply
	proto.Unmarshal(msg.Data, &rep)

	fmt.Printf("reply: %s\n", rep.Text)
}

// Output:
// reply: hello "joe"!