package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some data from the API"
	fmt.Fprintln(w, data)
}