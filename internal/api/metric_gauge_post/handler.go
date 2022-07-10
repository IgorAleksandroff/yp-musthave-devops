package metric_gauge_post

import (
	"net/http"
	"strconv"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics"
	"github.com/gorilla/mux"
)

type handler struct {
	metricsUC runtime_metrics.Usecase
}

func New(
	metricsUC runtime_metrics.Usecase,
) *handler {
	return &handler{
		metricsUC: metricsUC,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	metricName := vars["NAME"]
	metricValue, err := strconv.ParseFloat(vars["VALUE"], 64)
	if err != nil {
		http.Error(w, "can't parse a float64. internal error", http.StatusInternalServerError)
		return
	}

	h.metricsUC.SaveGaugeMetric(metricName, metricValue)

	w.WriteHeader(http.StatusOK)
}
