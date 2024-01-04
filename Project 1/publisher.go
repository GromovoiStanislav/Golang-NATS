package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Подключение к серверу NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Отправка сообщения каждую секунду
	for {
		message := "Привет, мир!"
		nc.Publish("topic", []byte(message))
		log.Printf("Отправлено сообщение: %s\n", message)
		time.Sleep(time.Second)
	}
}
