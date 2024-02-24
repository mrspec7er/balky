package user

import (
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

func HandlerConfig(q *amqp091.Channel, wg *sync.WaitGroup)  {
	a := &UserHandler{}

	a.CreateHandler(q, wg, "user.create", "user")
}