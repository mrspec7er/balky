package application

import (
	"os"
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

func HandlerConfig(q *amqp091.Channel, wg *sync.WaitGroup) {
	a := &ApplicationListener{}

	wg.Add(1)
	go a.CreateListener(q, wg, "app.create", "app.create."+os.Getenv("SERVER_ID"))
}
