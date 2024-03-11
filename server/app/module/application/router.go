package application

import (
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/balky/app/module/auth"
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

	wg.Add(1)
	go l.CreateReaction(q, wg, "reaction.create", "reaction.create."+os.Getenv("SERVER_ID"))
}

func RouteConfig(router chi.Router) {
	controller := &ApplicationController{}
	middleware := &auth.AuthMiddleware{}

	router.Get("/", controller.FindAll)
	router.With(middleware.Authorize("admin")).Post("/", controller.Create)
	router.With(middleware.Authorize("admin")).Delete("/", controller.Delete)

	router.Get("/contents/{id}", controller.FindAllContent)
	router.With(middleware.Authorize("admin")).Post("/contents", controller.CreateContent)
	router.With(middleware.Authorize("admin")).Delete("/contents", controller.DeleteContent)

	router.With(middleware.Authorize("admin")).Post("/reactions", controller.CreateReaction)
}
