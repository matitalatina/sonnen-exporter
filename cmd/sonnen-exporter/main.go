package main

import (
	"fmt"
	"log"
	"net/http"

	"mattianatali.it/sonnen-exporter/internal/metrics"
)

func main() {
	port := 2122
	http.HandleFunc("/metrics", metrics.HandleMetrics())
	log.Printf("Listening on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
