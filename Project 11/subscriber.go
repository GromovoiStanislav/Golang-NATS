package main

import (
	"log"
	"encoding/json"
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

	// Подписка на тему
	subject := "your_subject"
	nc.Subscribe(subject, func(msg *nats.Msg) {
		// Обработка полученного JSON-сообщения
		var receivedMessage Message
		err := json.Unmarshal(msg.Data, &receivedMessage)
		if err != nil {
			log.Println("Ошибка разбора JSON:", err)
			return
		}

		fmt.Printf("Получено JSON-сообщение: %+v", receivedMessage)
	})

	// Ожидание сообщений
	fmt.Println("Ожидание сообщений...")
	select {}
}
