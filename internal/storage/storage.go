package storage

import "fmt"

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

func (s *MemStorage) AllMetrics() string {
	var result string
	result += "Gauge metrics:\n"
	for n, v := range s.gaugeData {
		result += fmt.Sprintf("- %s = %f\n", n, v)
	}

	result += "Counter metrics:\n"
	for n, v := range s.counterData {
		result += fmt.Sprintf("- %s = %d\n", n, v)
	}

	return result
}
