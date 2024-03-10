package app

import (
	"context"
	"net/http"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/utility"
	"gorm.io/gorm"
)

type App struct {
	router       http.Handler
	dataListener func()
}

func New() *App {
	return &App{
		router:       loadRoutes(),
		dataListener: loadListeners,
	}
}

func (a *App) Start(ctx context.Context) error {
	utility.DBConnection()
	utility.AuthConfig()

	Migration(utility.DB)

	go a.dataListener()

	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(
		&model.Application{},
		&model.User{},
		&model.MasterReport{},
		&model.Attribute{},
		&model.Content{},
		&model.Logger{},
		&model.Reaction{},
	)
}
