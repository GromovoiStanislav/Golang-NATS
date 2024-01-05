package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
)

type msg struct {
	Client  string `json:"client"`
	Message string `json:"message"`
}

func main() {

	var channel string = "channel-test"
	var clientName string = generateClientName()

	c := connectNats()

	message := msg{
		Client:  clientName,
		Message: "The Dad is on",
	}


	sendMessage(c, channel, message)
	
}

func generateClientName() string {
	milisecondsNow := time.Now().UnixNano() / int64(time.Millisecond)

	clientName := "go-client-" + strconv.FormatInt(milisecondsNow, 10)

	fmt.Println("Client name:", clientName)

	return clientName
}



func connectNats() *nats.EncodedConn {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		fmt.Println("Error on conection")
	}

	fmt.Println("Connected on", nats.DefaultURL)

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	return c
}

func sendMessage(c *nats.EncodedConn, channel string, message msg) {

	err := c.Publish(channel, message)

	if err != nil {
		fmt.Println("Error on publish message")
	}

	fmt.Println("Send message to:", channel)
}