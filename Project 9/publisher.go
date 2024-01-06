package main

import (
	"log"
	"time"
	"strconv"

	"github.com/nats-io/nats.go"
)

func main() {
	// Подключение к серверу NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Отправка сообщений на тему "example"
	for i := 1; i <= 50; i++ {
		message := "Сообщение " + strconv.Itoa(i)
		nc.Publish("example", []byte(message))
		log.Printf("Отправлено сообщение: %s\n", message)
		time.Sleep(time.Second)
	}
}
