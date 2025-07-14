package main

import (
	"net/http"
	"fmt"

	"github.com/gorilla/mux"

	"web-scraper/crawler"
)

// func crawlForEngines(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	vin := vars["vin"]

// 	fmt.Fprintf(w, "Crawling for VIN: %s\n", vin)

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Engine found"))
// }

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{vin}", crawler.CrawlForEngines)

	fmt.Println("Server is running on port 8080/{vin}")
	http.ListenAndServe(":8080", router)
}

