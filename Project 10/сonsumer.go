package main

import (
	"context"
	"fmt"
	"os"
	"time"
	//"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()



	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	// connect to nats server
	nc, err := nats.Connect(url)
	if err != nil {
		fmt.Printf("Error connecting to NATS: %v\n", err)
		return
	}


	// create jetstream context from nats connection
	js, err := jetstream.New(nc)
	if err != nil {
		fmt.Printf("Error creating JetStream context: %v\n", err)
		return
	}

	

	//// Create new stream handle
	// _, err = js.CreateStream(ctx, jetstream.StreamConfig{
	// 	Name:     "foo",
	// 	Subjects: []string{"bar"},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }


	// pr get existing stream handle
	s, err := js.Stream(ctx, "foo")
	if err != nil {
		fmt.Printf("Error getting stream handle: %v\n", err)
		return
	}

	//// Create new consumer handle from a stream
	// cons, err := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
	// 	Durable:   "cons",
	// 	AckPolicy: jetstream.AckExplicitPolicy,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }


	// or retrieve consumer handle from a stream
	cons, err := s.Consumer(ctx, "cons")
	if err != nil {
		fmt.Printf("Error getting consumer handle: %v\n", err)
		return
	}


	// consume messages from the consumer in callback
	cc, err := cons.Consume(func(msg jetstream.Msg) {
		fmt.Println("Received jetstream message: ", string(msg.Data()))
		msg.Ack()
	})
	if err != nil {
		fmt.Printf("Error creating consumer: %v\n", err)
		return
	}
	defer cc.Stop()

	select {}
}
