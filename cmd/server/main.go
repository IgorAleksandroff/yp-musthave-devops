package main

import (
	"log"
	"net/http"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/handler/metricpost"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection/repository"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection/usecase"
)

func main() {
	metricsRepo := repository.New()
	metricsUC := usecase.New(metricsRepo)

	handler := metricpost.New(metricsUC)

	server := api.New()
	server.AddHandler(http.MethodPost, "/update/{TYPE}/{NAME}/{VALUE}", handler)

	log.Fatal(server.Run())
}
