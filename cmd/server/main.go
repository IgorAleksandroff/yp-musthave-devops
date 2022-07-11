package main

import (
	"log"
	"net/http"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/handlres/metricgaugepost"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection/entity"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection/repository"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection/usecase"
)

func main() {
	metricsStorage := map[string]entity.MetricGauge{}

	metricsRepo := repository.New(metricsStorage)
	metricsUC := usecase.New(metricsRepo)

	handler := metricgaugepost.New(metricsUC)

	server := api.New()
	server.AddHandler(http.MethodPost, "update/gauge/{NAME}/{VALUE}", handler)

	log.Fatal(server.Run())
}
