package storage

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
