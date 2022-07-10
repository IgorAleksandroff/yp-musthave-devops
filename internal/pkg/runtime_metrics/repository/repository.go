package repository

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics/entity"
)

type rep struct {
	storage map[string]entity.MetricGauge
}

func New(metricsStorage map[string]entity.MetricGauge) *rep {
	return &rep{storage: metricsStorage}
}
