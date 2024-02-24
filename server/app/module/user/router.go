package user

import (
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/rabbitmq/amqp091-go"
)

func HandlerConfig(q *amqp091.Channel, wg *sync.WaitGroup) {
	a := &UserHandler{}

	a.CreateHandler(q, wg, "user.create", "user")
}

func RouteConfig(router chi.Router) {
	controller := &UserController{}

	router.Get("/", controller.FindAllController)
}
