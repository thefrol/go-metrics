package storage

import (
	"fmt"
)

type gauge float64
type Counter int64

type Metrics interface {
	UpdateCounter(name string, value int64)
	UpdateGauge(name string, value float64)
	GetValue(t string, name string) string
}

type MemStorage struct {
	gaugeData   map[string]gauge
	counterData map[string]Counter
}

// NewMemStorage returns a new in memory storage instance.
func NewMemStorage() *MemStorage {

	return &MemStorage{
		gaugeData:   make(map[string]gauge),
		counterData: make(map[string]Counter),
	}
}

func (s *MemStorage) UpdateCounter(name string, value int64) {
	s.counterData[name] += Counter(value)
}

func (s *MemStorage) UpdateGauge(name string, value float64) {
	s.gaugeData[name] = gauge(value)
}

func (s *MemStorage) GetValue(t string, name string) string {
	var v string
	//	statusCode := http.StatusOK
	if val, ok := s.gaugeData[name]; ok && t == "gauge" {
		v = fmt.Sprint(val)
	} else if val, ok := s.counterData[name]; ok && t == "counter" {
		v = fmt.Sprint(val)
	}
	return v
}

func (s *MemStorage) AllMetrics() string {
	var result string
	//TODO
	//Формирование строкового представления метрик тоже зона ответственности хендлера, Storage должен возвращать сами данные.
	//	Можно вернуть сами мапы, предварительно скопировав их, либо создать отдельные структуры данных и возвращать список этих структур
	//result += "Gauge metrics:\n"
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

//func (s *MemStorage) GetAll() {
//	counter := s.counterData
//	for key, _ := range counter {
//		fmt.Printf(key)
//	}
//	k := s.counterData[]
//i := 0
////	v :=
//for _, v = range s.counterData {
//	counter[i] = v
//	i++
//
//	println(k, v)
//}
//rv := make([]gauge, len(s.gaugeData))
//
//i := 0
//for _, v := range s.gaugeData {
//	rv[i] = v
//	i++
//}
//	println(rv)
//return rv
//println(counter)
//return counter
//}

//func (s *MemStorage) Get()  {
//	record := range s.counterData
//		{
//		println record
//	}
//	return record
//}
