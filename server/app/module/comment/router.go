package comment

import (
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/balky/app/module/auth"
	"github.com/rabbitmq/amqp091-go"
)

func HandlerConfig(q *amqp091.Channel, wg *sync.WaitGroup) {
	l := &CommentListener{}

	wg.Add(1)
	go l.Create(q, wg, "comment.create", "comment.create."+os.Getenv("SERVER_ID"))

	wg.Add(1)
	go l.DeleteListener(q, wg, "comment.delete", "comment.delete."+os.Getenv("SERVER_ID"))
}

func RouteConfig(router chi.Router) {
	controller := &CommentController{}
	middleware := &auth.AuthMiddleware{}

	router.Get("/{appId}", controller.FindAll)
	router.With(middleware.Authorize("admin")).Post("/", controller.Create)
	router.With(middleware.Authorize("admin")).Delete("/", controller.Delete)
}
