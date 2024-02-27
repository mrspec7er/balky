package master

import (
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/rabbitmq/amqp091-go"
)

func HandlerConfig(q *amqp091.Channel, wg *sync.WaitGroup) {
	l := &MasterReportListener{}

	wg.Add(1)
	go l.Create(q, wg, "master.create", "master.create."+os.Getenv("SERVER_ID"))

	wg.Add(1)
	go l.DeleteListener(q, wg, "master.delete", "master.delete."+os.Getenv("SERVER_ID"))
}

func RouteConfig(router chi.Router) {
	controller := &MasterReportController{}

	router.Get("/", controller.FindAll)
	router.Post("/", controller.Create)
	router.Delete("/", controller.Delete)
}
