package app

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/mrspec7er/balky/app/module/application"
	"github.com/mrspec7er/balky/app/module/user"
	"github.com/rabbitmq/amqp091-go"
)

func loadListeners() {
	conn, err := amqp091.Dial(os.Getenv("MESSAGE_BROKER_URI"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	wg := sync.WaitGroup{}

	listenersConfig(ch, &wg)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	fmt.Println("Interrupt signal received. Closing consumers...")
	wg.Wait()
	fmt.Println("All consumers closed.")
}

func listenersConfig(ch *amqp091.Channel, wg *sync.WaitGroup) {
	wg.Add(1)
	go user.HandlerConfig(ch, wg)

	wg.Add(1)
	go application.HandlerConfig(ch, wg)
}
