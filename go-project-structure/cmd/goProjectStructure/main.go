package main

import (
	"fmt"
	"net/http"

	"github.com/ShakthiW/goProjectStructure/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server is running on port http://localhost:%d\n", port)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		panic(err)
	}
}