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

	// Отправка запроса и ожидание ответа
	response, err := nc.Request("request", []byte("Пример запроса"), 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Получен ответ: %s\n", string(response.Data))
}
