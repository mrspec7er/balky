package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/mrspec7er/balky/app"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	server := app.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)	
	defer cancel()

	err := server.Start(ctx)
	if err != nil {
		panic(err)
	}
}
