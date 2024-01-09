package main

import (
	"encoding/json"
	"log"
	"fmt"

	"github.com/nats-io/nats.go"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	// Подключение к серверу NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Создание JSON-сообщения
	message := Message{Text: "Hello NATS!"}
	payload, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	// Отправка JSON-сообщения
	subject := "your_subject"
	err = nc.Publish(subject, payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Отправлено JSON-сообщение: %s", payload)
}
