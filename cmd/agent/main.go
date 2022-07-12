package main

import (
	"time"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/services/devopsserver"
	runtimeMetricsRepository "github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtimemetrics/repository"
	runtimeMetricsUsecase "github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtimemetrics/usecase"
)

func main() {
	pollInterval := time.NewTicker(2 * time.Second)
	reportInterval := time.NewTicker(10 * time.Second)

	client := devopsserver.NewClient()
	runtimeMetricsRepo := runtimeMetricsRepository.New()
	runtimeMetricsUC := runtimeMetricsUsecase.New(runtimeMetricsRepo, client)

	for {
		select {
		case <-pollInterval.C:
			runtimeMetricsUC.UpdateMetrics()
		case <-reportInterval.C:
			runtimeMetricsUC.SendMetrics()
		}
	}
}
