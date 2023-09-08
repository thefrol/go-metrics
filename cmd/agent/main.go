package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Config struct {
	address        string `env:"ADDRESS"`
	reportInterval int    `env:"REPORT_INTERVAL"`
	pollInterval   int    `env:"POLL_INTERVAL"`
}

var cfg Config
var valuesGauge = map[string]float64{}
var pollCount uint64

//var pollInterval int
//var reportInterval int
//var flagRunAddr string

func getMetrics() {
	var rtm runtime.MemStats
	pollCount += 1
	// Read full mem stats
	runtime.ReadMemStats(&rtm)

	valuesGauge["Alloc"] = float64(rtm.Alloc)
	valuesGauge["BuckHashSys"] = float64(rtm.BuckHashSys)
	valuesGauge["Frees"] = float64(rtm.Frees)
	valuesGauge["GCCPUFraction"] = float64(rtm.GCCPUFraction)
	valuesGauge["HeapAlloc"] = float64(rtm.HeapAlloc)
	valuesGauge["HeapIdle"] = float64(rtm.HeapIdle)
	valuesGauge["HeapInuse"] = float64(rtm.HeapInuse)
	valuesGauge["HeapObjects"] = float64(rtm.HeapObjects)
	valuesGauge["HeapReleased"] = float64(rtm.HeapReleased)
	valuesGauge["HeapSys"] = float64(rtm.HeapSys)
	valuesGauge["LastGC"] = float64(rtm.LastGC)
	valuesGauge["Lookups"] = float64(rtm.Lookups)
	valuesGauge["MCacheInuse"] = float64(rtm.MCacheInuse)
	valuesGauge["MCacheSys"] = float64(rtm.MCacheSys)
	valuesGauge["MSpanInuse"] = float64(rtm.MSpanInuse)
	valuesGauge["MSpanSys"] = float64(rtm.MSpanSys)
	valuesGauge["Mallocs"] = float64(rtm.Mallocs)
	valuesGauge["NextGC"] = float64(rtm.NextGC)
	valuesGauge["NumForcedGC"] = float64(rtm.NumForcedGC)
	valuesGauge["NumGC"] = float64(rtm.NumGC)
	valuesGauge["OtherSys"] = float64(rtm.OtherSys)
	valuesGauge["PauseTotalNs"] = float64(rtm.PauseTotalNs)
	valuesGauge["StackInuse"] = float64(rtm.StackInuse)
	valuesGauge["StackSys"] = float64(rtm.StackSys)
	valuesGauge["Sys"] = float64(rtm.Sys)
	valuesGauge["TotalAlloc"] = float64(rtm.TotalAlloc)

}

func postQueries() {
	for k, v := range valuesGauge {
		post("gauge", k, strconv.FormatFloat(v, 'f', -1, 64))
	}
	post("counter", "PollCount", strconv.FormatUint(pollCount, 10))
	post("gauge", "RandomValue", strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
	pollCount = 0
}

func post(t string, name string, value string) {
	bodyReader := bytes.NewReader([]byte{})

	// We can set the content type here
	fmt.Println("Running server on", cfg.address)
	resp, err := http.Post(fmt.Sprintf("http://%s/update/%s/%s/%s", cfg.address, t, name, value), "text/plain", bodyReader)
	//resp, err := http.Post(url, "text/plain", bodyReader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("POST:", resp.Request)
}

func runParameters() error {
	flag.StringVar(&cfg.address, "a", "localhost:8080", "address and port to run server")
	flag.IntVar(&cfg.reportInterval, "r", 10, "frequency of sending metrics to the server")
	flag.IntVar(&cfg.pollInterval, "p", 2, "frequency of polling metrics")
	flag.Parse()

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		cfg.address = envRunAddr
	}
	if envRunAddr := os.Getenv("REPORT_INTERVAL"); envRunAddr != "" {
		cfg.reportInterval, _ = strconv.Atoi(envRunAddr)
	}
	if envRunAddr := os.Getenv("POLL_INTERVAL"); envRunAddr != "" {
		cfg.pollInterval, _ = strconv.Atoi(envRunAddr)
	}

	return nil
}
func main() {
	err := runParameters()
	if err != nil {
		log.Fatal(err)
	}
	pollTicker := time.NewTicker(time.Duration(cfg.pollInterval) * time.Second)
	defer pollTicker.Stop()
	reportTicker := time.NewTicker(time.Duration(cfg.reportInterval) * time.Second)
	defer reportTicker.Stop()

	for {
		select {
		case <-pollTicker.C:
			getMetrics()
			b, _ := json.Marshal(valuesGauge)
			fmt.Println(string(b))
		case <-reportTicker.C:
			postQueries()
		}
	}
}
