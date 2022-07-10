package main

import (
	"log"
	"net/http"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/metric_gauge_post"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics/entity"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics/repository"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics/usecase"
)

func main() {
	metricsStorage := map[string]entity.MetricGauge{}

	metricsRepo := repository.New(metricsStorage)
	metricsUC := usecase.New(metricsRepo)

	handler := metric_gauge_post.New(metricsUC)

	server := api.New()
	server.AddHandler(http.MethodPost, "update/gauge/{NAME}/{VALUE}", handler)

	log.Fatal(server.Run())
}
