package metrics

import "runtime"

var ValuesGauge = map[string]float64{}

var PollCount uint64

func GetMetrics() {
	var rtm runtime.MemStats
	PollCount += 1
	// Read full mem stats
	runtime.ReadMemStats(&rtm)

	ValuesGauge["Alloc"] = float64(rtm.Alloc)
	ValuesGauge["BuckHashSys"] = float64(rtm.BuckHashSys)
	ValuesGauge["Frees"] = float64(rtm.Frees)
	ValuesGauge["GCCPUFraction"] = float64(rtm.GCCPUFraction)
	ValuesGauge["HeapAlloc"] = float64(rtm.HeapAlloc)
	ValuesGauge["HeapIdle"] = float64(rtm.HeapIdle)
	ValuesGauge["HeapInuse"] = float64(rtm.HeapInuse)
	ValuesGauge["HeapObjects"] = float64(rtm.HeapObjects)
	ValuesGauge["HeapReleased"] = float64(rtm.HeapReleased)
	ValuesGauge["HeapSys"] = float64(rtm.HeapSys)
	ValuesGauge["LastGC"] = float64(rtm.LastGC)
	ValuesGauge["Lookups"] = float64(rtm.Lookups)
	ValuesGauge["MCacheInuse"] = float64(rtm.MCacheInuse)
	ValuesGauge["MCacheSys"] = float64(rtm.MCacheSys)
	ValuesGauge["MSpanInuse"] = float64(rtm.MSpanInuse)
	ValuesGauge["MSpanSys"] = float64(rtm.MSpanSys)
	ValuesGauge["Mallocs"] = float64(rtm.Mallocs)
	ValuesGauge["NextGC"] = float64(rtm.NextGC)
	ValuesGauge["NumForcedGC"] = float64(rtm.NumForcedGC)
	ValuesGauge["NumGC"] = float64(rtm.NumGC)
	ValuesGauge["OtherSys"] = float64(rtm.OtherSys)
	ValuesGauge["PauseTotalNs"] = float64(rtm.PauseTotalNs)
	ValuesGauge["StackInuse"] = float64(rtm.StackInuse)
	ValuesGauge["StackSys"] = float64(rtm.StackSys)
	ValuesGauge["Sys"] = float64(rtm.Sys)
	ValuesGauge["TotalAlloc"] = float64(rtm.TotalAlloc)

}
