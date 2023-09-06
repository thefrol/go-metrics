package main

import (
	"github.com/Sofja96/go-metrics.git/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

func metrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	s := storage.New()
	sliceURL := strings.Split(r.URL.Path, "/")

	if len(sliceURL) != 5 || sliceURL[1] != "update" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	metricsType := sliceURL[2]
	metricsName := sliceURL[3]
	metricsValue := sliceURL[4]
	if metricsType == "counter" {
		if value, err := strconv.ParseInt(metricsValue, 10, 64); err == nil {
			s.UpdateCounter(metricsName, value)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else if metricsType == "gauge" {
		if value, err := strconv.ParseFloat(metricsValue, 64); err == nil {
			s.UpdateGauge(metricsName, value)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, metrics)
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
