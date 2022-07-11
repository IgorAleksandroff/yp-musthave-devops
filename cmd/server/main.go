package main

import (
	"log"
	"net/http"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/handlres/metricGaugePost"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricsCollection/entity"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricsCollection/repository"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricsCollection/usecase"
)

func main() {
	metricsStorage := map[string]entity.MetricGauge{}

	metricsRepo := repository.New(metricsStorage)
	metricsUC := usecase.New(metricsRepo)

	handler := metricGaugePost.New(metricsUC)

	server := api.New()
	server.AddHandler(http.MethodPost, "update/gauge/{NAME}/{VALUE}", handler)

	log.Fatal(server.Run())
}
