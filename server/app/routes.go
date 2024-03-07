package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrspec7er/balky/app/module/application"
	"github.com/mrspec7er/balky/app/module/auth"
	"github.com/mrspec7er/balky/app/module/logger"
	"github.com/mrspec7er/balky/app/module/master"
	"github.com/mrspec7er/balky/app/module/user"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello There!"))
	})

	router.Route("/users", user.RouteConfig)
	router.Route("/masters", master.RouteConfig)
	router.Route("/loggers", logger.RouteConfig)
	router.Route("/auth", auth.RouteConfig)
	router.Route("/applications", application.RouteConfig)

	return router
}
