package handlers

import (
	"fmt"
	"github.com/Sofja96/go-metrics.git/internal/storage"
	_ "github.com/Sofja96/go-metrics.git/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// разрешаем только POST-запросы
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

func AllMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// разрешаем только GET-запросы
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	s := storage.New()
	//data := ioutil.ReadAll(w.Body)
	fmt.Println(string(s.AllMetrics()))
	metric := s.AllMetrics()
	fmt.Println(metric)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
}
