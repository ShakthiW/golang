package main

import (
	"context"
	"fmt"
	"orderapi/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}