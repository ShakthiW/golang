package main

import (
	"context"
	"fmt"
	"orderapi/application"
	"os"
	"os/signal"
)

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
