package main

import (
	"fmt"
	"os"
	"time"
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


	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Responding to a request message
	nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 42"))
	})

	// Requests
	go func(){
		response, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)
		if err != nil {
			log.Println(err)
		}
		if response != nil {
			fmt.Printf("Received a answer: %s\n", string(response.Data))
		}
	}()

	// Requests
	go func(){
		response, err := nc.Request("request", nil, 10*time.Millisecond)
		if err != nil {
			log.Println(err)
		}
		if response != nil {
			fmt.Printf("Received a answer: %s\n", string(response.Data))
		}
	}()
	
	
	select {}
}