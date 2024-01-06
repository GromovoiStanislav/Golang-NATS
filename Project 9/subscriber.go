package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Подключение к серверу NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Подписка на тему "example" с использованием Queue Group "workers"
	_, err = nc.QueueSubscribe("example", "workers", func(msg *nats.Msg) {
		log.Printf("Получено сообщение: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	select {}
}