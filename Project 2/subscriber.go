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

	// Подписка на темы с использованием маски
	nc.Subscribe("foo.*.*", func(msg *nats.Msg) {
		log.Printf("Получено сообщение по маске: %s\n", string(msg.Data))
	})

	select {}
}
