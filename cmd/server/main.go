package main

import (
	"log"
	"net/http"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/handlres/metric_gauge_post"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metrics_collection/entity"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metrics_collection/repository"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metrics_collection/usecase"
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
