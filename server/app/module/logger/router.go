package logger

import (
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/rabbitmq/amqp091-go"
)

func HandlerConfig(q *amqp091.Channel, wg *sync.WaitGroup) {
	l := &LoggerListener{}

	wg.Add(1)
	go l.Create(q, wg, "logger", "logger."+os.Getenv("SERVER_ID"))
}

func RouteConfig(router chi.Router) {
	controller := &LoggerController{}

	router.Get("/", controller.FindAll)
}
