package metricpost

import (
	"net/http"
	"strconv"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection/entity"
	"github.com/gorilla/mux"
)

type handler struct {
	metricsUC metricscollection.Usecase
}

func New(
	metricsUC metricscollection.Usecase,
) *handler {
	return &handler{
		metricsUC: metricsUC,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	metricType := vars["TYPE"]
	metricName := vars["NAME"]

	switch metricType {
	case entity.CounterTypeMetric:
		counterValue, err := strconv.ParseInt(vars["VALUE"], 10, 64)
		if err != nil {
			http.Error(w, "can't parse a int64. internal error", http.StatusBadRequest)
			return
		}

		h.metricsUC.SaveCounterMetric(metricName, counterValue)
	case entity.GaugeTypeMetric:
		gaugeValue, err := strconv.ParseFloat(vars["VALUE"], 64)
		if err != nil {
			http.Error(w, "can't parse a float64. internal error", http.StatusBadRequest)
			return
		}

		h.metricsUC.SaveGaugeMetric(metricName, gaugeValue)
	default:
		http.Error(w, "unknown handler", http.StatusNotImplemented)
		return
	}

	w.WriteHeader(http.StatusOK)
}
