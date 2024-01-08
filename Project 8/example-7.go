package main

import (
	"fmt"
	"os"


	"github.com/nats-io/nats.go"
)

func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()

	type person struct {
		Name     string
		Address  string
		Age      int
	}

	recvCh := make(chan *person)
	ec.BindRecvChan("hello", recvCh)

	sendCh := make(chan *person)
	ec.BindSendChan("hello", sendCh)

	me := &person{Name: "derek", Age: 22, Address: "140 New Montgomery Street"}

	// Send via Go channels
	sendCh <- me

	// Receive via Go channels
	who := <- recvCh
	fmt.Printf("Received a person: %+v\n", who)
}