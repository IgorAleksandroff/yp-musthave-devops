package main

import (
	"time"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/services/devopsServer"
	runtimeMetricsRepository "github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtimeMetrics/repository"
	runtimeMetricsUsecase "github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtimeMetrics/usecase"
)

func main() {
	pollInterval := time.NewTicker(2 * time.Second)
	reportInterval := time.NewTicker(10 * time.Second)

	client := devopsServer.NewClient()
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
