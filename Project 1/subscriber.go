package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func main() {
	// Подключение к серверу NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Создание подписчика на тему "topic"
	_, err = nc.Subscribe("topic", func(msg *nats.Msg) {
		log.Printf("Получено сообщение: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// Ожидание сигнала для завершения программы
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel
}
