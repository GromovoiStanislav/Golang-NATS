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

	c := connectNats()

	subscribeChannel(c, channel)
	continueListen()
}

func generateClientName() string {
	milisecondsNow := time.Now().UnixNano() / int64(time.Millisecond)

	clientName := "go-client-" + strconv.FormatInt(milisecondsNow, 10)

	fmt.Println("Client name:", clientName)

	return clientName
}

func continueListen() {
	fmt.Println("Listen")
	select {}
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

func subscribeChannel(c *nats.EncodedConn, channel string) {

	_, err := c.Subscribe(channel, func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	if err != nil {
		fmt.Println("Error on subcribe on a chanel")
	}

	fmt.Println("Subscribe successfuly on", channel)
}