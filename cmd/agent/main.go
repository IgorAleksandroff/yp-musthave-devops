package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
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

func main() {
	memMetrics := runtime.MemStats{}
	runtime.ReadMemStats(&memMetrics)

	var curMetrics []metric
	curMetrics = []metric{
		{"Alloc", gauge(float64(memMetrics.Alloc))},
		{"Frees", gauge(float64(memMetrics.Frees))},
	}

	fmt.Println(curMetrics)

	// client
	client := &http.Client{}

	hostName := "127.0.0.1"
	endpoint := fmt.Sprintf("http://%s:8080/update/%s/%s/%v/", hostName, curMetrics[0].value.getType(), curMetrics[0].name, curMetrics[0].value)
	fmt.Println(endpoint)

	request, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	request.Header.Add("Content-Type", "plain")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Статус-код ", response.Status)
}
