package main

import (
	"fmt"
	"log"
	"context"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	
	// Подключение к серверу NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()


	// Создание JetStream контекста из соединения NATS
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	// Отправка сообщения в JetStream стрим
	subject := "bar" // Субъект сообщения
	message := []byte("Hello, JetStream!")

	// Публикация сообщения в JetStream
	_, err = js.Publish(ctx, subject, message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Сообщение успешно опубликовано в JetStream стрим.")
}