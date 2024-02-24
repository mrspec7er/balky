package application

import (
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

func HandlerConfig(q *amqp091.Channel, wg *sync.WaitGroup) {
	a := &ApplicationHandler{}

	a.CreateHandler(q, wg, "app.create", "application")
}