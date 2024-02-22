package app

import (
	"context"
	"net/http"
	"os"

	"github.com/mrspec7er/balky/app/utils"
)

type App struct {
	router http.Handler
	handler func()
}

func New() *App {
	return &App{
		router: loadRoutes(),
		handler: loadHandlers,
	}
}

func (a *App) Start(ctx context.Context) error {
	utils.DBConnection()

	a.handler()
	server := &http.Server{
		Addr: os.Getenv("PORT"),
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}