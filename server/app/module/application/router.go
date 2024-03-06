package application

import (
	"os"
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

func HandlerConfig(q *amqp091.Channel, wg *sync.WaitGroup) {
	l := &ApplicationListener{}

	wg.Add(1)
	go l.Create(q, wg, "app.create", "app.create."+os.Getenv("SERVER_ID"))

	wg.Add(1)
	go l.Delete(q, wg, "app.delete", "app.delete."+os.Getenv("SERVER_ID"))

	wg.Add(1)
	go l.CreateContent(q, wg, "content.create", "content.create."+os.Getenv("SERVER_ID"))

	wg.Add(1)
	go l.DeleteContent(q, wg, "content.delete", "content.delete."+os.Getenv("SERVER_ID"))
}
