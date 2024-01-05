package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)
	defer nc.Drain()

	// Simple Sync Subscriber
	sub, _ := nc.SubscribeSync("foo")


	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	msg, _ := sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("Received a message: %s\n", string(msg.Data))

	// Unsubscribe
	sub.Unsubscribe()
}