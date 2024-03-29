package master

import (
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/balky/app/module/auth"
	"github.com/rabbitmq/amqp091-go"
)

func HandlerConfig(q *amqp091.Channel, wg *sync.WaitGroup) {
	l := &MasterReportListener{}

	wg.Add(1)
	go l.Create(q, wg, "master.create", "master.create."+os.Getenv("SERVER_ID"))

	wg.Add(1)
	go l.Delete(q, wg, "master.delete", "master.delete."+os.Getenv("SERVER_ID"))

	wg.Add(1)
	go l.CreateAttribute(q, wg, "attribute.create", "attribute.create."+os.Getenv("SERVER_ID"))

	wg.Add(1)
	go l.DeleteAttribute(q, wg, "attribute.delete", "attribute.delete."+os.Getenv("SERVER_ID"))
}

func RouteConfig(router chi.Router) {
	controller := &MasterReportController{}
	middleware := &auth.AuthMiddleware{}

	router.Get("/", controller.FindAll)
	router.With(middleware.Authorize("admin")).Post("/", controller.Create)
	router.With(middleware.Authorize("admin")).Delete("/", controller.Delete)

	router.Get("/attributes", controller.FindAllAttribute)
	router.With(middleware.Authorize("admin")).Post("/attributes", controller.CreateAttribute)
	router.With(middleware.Authorize("admin")).Delete("/attributes", controller.DeleteAttribute)
}
