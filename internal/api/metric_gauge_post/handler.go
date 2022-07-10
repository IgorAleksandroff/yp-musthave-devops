package metric_gauge_post

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type handler struct {
	metricsUC metrics.Usecase
}

func New(
	metricsUC metrics.Usecase,
) *handler {
	return &handler{
		metricsUC: metricsUC,
	}
}

func (h *handler) GaugeHandler(w http.ResponseWriter, r *http.Request) {
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
