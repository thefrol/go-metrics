package storage

import (
	"fmt"
	"net/http"
)

type gauge float64
type counter int64

type MemStorage struct {
	gaugeData   map[string]gauge
	counterData map[string]counter
}

// New returns a new in memory storage instance.
func New() *MemStorage {

	return &MemStorage{
		gaugeData:   make(map[string]gauge),
		counterData: make(map[string]counter),
	}
}

func (s *MemStorage) UpdateCounter(name string, value int64) {
	s.counterData[name] += counter(value)
}

func (s *MemStorage) UpdateGauge(name string, value float64) {
	s.gaugeData[name] = gauge(value)
}

func (s *MemStorage) GetValue(t string, name string) (string, int) {
	var v string
	statusCode := http.StatusOK
	if val, ok := s.gaugeData[name]; ok && t == "gauge" {
		v = fmt.Sprint(val)
	} else if val, ok := s.counterData[name]; ok && t == "counter" {
		v = fmt.Sprint(val)
	} else {
		statusCode = http.StatusNotFound
	}
	return v, statusCode
}

func (s *MemStorage) AllMetrics() string {
	var result string
	result += "Gauge metrics:\n"
	for name, value := range s.gaugeData {
		result += fmt.Sprintf("- %s = %f\n", name, value)
	}

	result += "Counter metrics:\n"
	for name, value := range s.counterData {
		result += fmt.Sprintf("- %s = %d\n", name, value)
	}

	return result
}
