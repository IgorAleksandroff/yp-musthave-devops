package main

import (
	"time"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/services/devops_server"
	runtimeMetricsRepository "github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics/repository"
	runtimeMetricsUsecase "github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics/usecase"
)

func main() {
	pollInterval := time.NewTicker(2 * time.Second)
	reportInterval := time.NewTicker(10 * time.Second)

	client := devops_server.NewClient()
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
