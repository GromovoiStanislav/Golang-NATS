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

	// connect to nats server
	nc, _ := nats.Connect(url)
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()


	// Simple Async Subscriber
	sub, _ := ec.Subscribe("foo", func(s string) {
		fmt.Printf("Received a message: %s\n", s)
	})
	// Simple Publisher
	ec.Publish("foo", "Hello World")


	// EncodedConn can Publish any raw Go type using the registered Encoder
	type person struct {
		Name     string
		Address  string
		Age      int
	}
	// Go type Subscriber
	ec.Subscribe("hello", func(p *person) {
		fmt.Printf("Received a person: %+v\n", p)
	})
	me := &person{Name: "derek", Age: 22, Address: "140 New Montgomery Street, San Francisco, CA"}
	// Go type Publisher
	ec.Publish("hello", me)



	// Replying
	ec.Subscribe("help", func(subj, reply string, msg string) {
		ec.Publish(reply, "I can help!")
	})
	// Requests
	var response string
	err := ec.Request("help", "help me", &response, 10*time.Millisecond)
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	}
	fmt.Printf("Received a answer: %s\n", string(response))
	


	// Unsubscribe
	sub.Unsubscribe()

	select {}
}
