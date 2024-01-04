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

	// Отправка сообщений на темы с использованием маски
	for i := 1; i <= 3; i++ {
		message := "Сообщение " + strconv.Itoa(i)
		topic := "foo.bar." + strconv.Itoa(i)
		nc.Publish(topic, []byte(message))
		log.Printf("Отправлено сообщение по теме %s: %s\n", topic, message)
		time.Sleep(time.Second)
	}
}
