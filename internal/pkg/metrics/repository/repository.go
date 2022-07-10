package repository

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metrics/entity"
)

type DB struct {
	storage map[string]entity.MetricGauge
}

type rep struct {
	metricsStorage DB
}

func New(metricsStorage DB) *rep {
	return &rep{metricsStorage: metricsStorage}
}
