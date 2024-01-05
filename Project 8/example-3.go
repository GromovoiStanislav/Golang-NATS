package main

import (
	"fmt"
	"os"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)
	defer nc.Drain()

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe("foo", ch)
	if err != nil {
		log.Fatal(err)
	}
	

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	msg := <- ch
	fmt.Printf("Received a message: %s\n", string(msg.Data))

	// Unsubscribe
	sub.Unsubscribe()
}