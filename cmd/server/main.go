package main

import (
	"log"
	"net/http"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/handler/metriccounterpost"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/api/handler/metricgaugepost"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection/repository"
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection/usecase"
)

func main() {
	metricsRepo := repository.New()
	metricsUC := usecase.New(metricsRepo)

	gaugeHandler := metricgaugepost.New(metricsUC)
	countHandler := metriccounterpost.New(metricsUC)

	server := api.New()
	server.AddHandler(http.MethodPost, "/update/gauge/{NAME}/{VALUE}", gaugeHandler)
	server.AddHandler(http.MethodPost, "/update/counter/{NAME}/{VALUE}", countHandler)

	log.Fatal(server.Run())
}
