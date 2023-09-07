package main

import (
	"github.com/Sofja96/go-metrics.git/internal/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handlers.Webhook)
	//mux.HandleFunc(`/all`, handlers.AllMetrics)
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
