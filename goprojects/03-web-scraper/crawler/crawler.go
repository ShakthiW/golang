package crawler

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

func CrawlForEngines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vin := vars["vin"]

	// Create new Colly Collector
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.MaxDepth(100),
	)

	// Hnadle the POST call with the vin
	url := "https://www.hollanderparts.com/Home"
	c.PostMultipart(
		url, map[string][]byte{
			"hdnVIN": []byte(vin),
		},
	)

	fmt.Fprintf(w, "Crawling for VIN: %s\n", vin)
}