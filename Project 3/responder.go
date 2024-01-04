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

	// Определение обработчика для запросов
	nc.Subscribe("request", func(msg *nats.Msg) {
		// Обработка запроса и отправка ответа
		response := "Ответ на ваш запрос"
		nc.Publish(msg.Reply, []byte(response))
		log.Printf("Получен запрос: %s, Отправлен ответ: %s\n", string(msg.Data), response)
	})
	select {}
}
