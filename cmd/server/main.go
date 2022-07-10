package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type metricGauge struct {
	typeMetric string
	value      float64
}

var metricsStorage = map[string]metricGauge{}

func GaugeHandler(w http.ResponseWriter, r *http.Request) {
	typeMetric := "gauge"
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	metricName := vars["NAME"]
	metricValue, err := strconv.ParseFloat(vars["VALUE"], 64)
	if err != nil {
		http.Error(w, "can't parse a float64. internal error", http.StatusInternalServerError)
		return
	}

	metricsStorage[metricName] = metricGauge{
		typeMetric: typeMetric,
		value:      metricValue,
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("update/gauge/{NAME}/{VALUE}", GaugeHandler).Methods(http.MethodPost)
	http.Handle("/", router)

	log.Println("start server")

	log.Fatal(http.ListenAndServe(":8080", router))
}
