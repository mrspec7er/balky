package user

import (
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/rabbitmq/amqp091-go"
)

func HandlerConfig(q *amqp091.Channel, wg *sync.WaitGroup) {
	l := &UserListener{}

	wg.Add(1)
	go l.Create(q, wg, "user.create", "user.create."+os.Getenv("SERVER_ID"))

	wg.Add(1)
	go l.DeleteListener(q, wg, "user.delete", "user.delete."+os.Getenv("SERVER_ID"))
}

func RouteConfig(router chi.Router) {
	controller := &UserController{}

	router.Get("/", controller.FindAll)
	router.Post("/", controller.Create)
	router.Delete("/", controller.Delete)
}
