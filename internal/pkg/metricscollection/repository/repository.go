package repository

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/metricscollection/entity"
)

type rep struct {
	storage map[string]entity.MetricGauge
}

func New(metricsStorage map[string]entity.MetricGauge) *rep {
	return &rep{storage: metricsStorage}
}
