package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
)

type gauge float64
type counter int64

type getter interface {
	getType() string
}

type metric struct {
	name  string
	value getter
}

func (gauge) getType() string {
	return "gauge"
}

func (counter) getType() string {
	return "counter"
}

var curMetrics []metric
var pollCount int64
var randomValue int64

func updateMetrics() {
	pollCount++
	randomValue = int64(rand.Int())
	memMetrics := runtime.MemStats{}
	runtime.ReadMemStats(&memMetrics)

	curMetrics = []metric{
		{"Alloc", gauge(float64(memMetrics.Alloc))},
		{"BuckHashSys", gauge(float64(memMetrics.BuckHashSys))},
		{"Frees", gauge(float64(memMetrics.Frees))},
		{"GCCPUFraction", gauge(memMetrics.GCCPUFraction)},
		{"GCSys", gauge(float64(memMetrics.GCSys))},
		{"HeapAlloc", gauge(float64(memMetrics.HeapAlloc))},
		{"HeapIdle", gauge(float64(memMetrics.HeapIdle))},
		{"HeapInuse", gauge(float64(memMetrics.HeapInuse))},
		{"HeapObjects", gauge(float64(memMetrics.HeapObjects))},
		{"HeapReleased", gauge(float64(memMetrics.HeapReleased))},
		{"HeapSys", gauge(float64(memMetrics.HeapSys))},
		{"LastGC", gauge(float64(memMetrics.LastGC))},
		{"Lookups", gauge(float64(memMetrics.Lookups))},
		{"MCacheInuse", gauge(float64(memMetrics.MCacheInuse))},
		{"MCacheSys", gauge(float64(memMetrics.MCacheSys))},
		{"MSpanInuse", gauge(float64(memMetrics.MSpanInuse))},
		{"MSpanSys", gauge(float64(memMetrics.MSpanSys))},
		{"Mallocs", gauge(float64(memMetrics.Mallocs))},
		{"NextGC", gauge(float64(memMetrics.NextGC))},
		{"NumForcedGC", gauge(float64(memMetrics.NumForcedGC))},
		{"NumGC", gauge(float64(memMetrics.NumGC))},
		{"OtherSys", gauge(float64(memMetrics.OtherSys))},
		{"PauseTotalNs", gauge(float64(memMetrics.PauseTotalNs))},
		{"StackInuse", gauge(float64(memMetrics.StackInuse))},
		{"StackSys", gauge(float64(memMetrics.StackSys))},
		{"Sys", gauge(float64(memMetrics.Sys))},
		{"TotalAlloc", gauge(float64(memMetrics.TotalAlloc))},
		{"PollCount", counter(pollCount)},
		{"RandomValue", counter(randomValue)},
	}
}

func sendMetric(client *http.Client, endpoint string) error {
	request, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "plain")

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

func sendMetrics(client *http.Client) error {
	hostName := "127.0.0.1"
	for _, curMetric := range curMetrics {
		endpoint := fmt.Sprintf("http://%s:8080/update/%s/%s/%v/", hostName, curMetric.value.getType(), curMetric.name, curMetric.value)

		if err := sendMetric(client, endpoint); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// client
	client := &http.Client{}

	pollInterval := time.NewTicker(2 * time.Second)
	reportInterval := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-pollInterval.C:
			updateMetrics()
		case <-reportInterval.C:
			if err := sendMetrics(client); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
