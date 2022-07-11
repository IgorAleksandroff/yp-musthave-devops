package repository

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtimeMetrics/entity"
)

type rep struct {
	storage map[string]entity.Metric
}

func New() *rep {
	return &rep{storage: make(map[string]entity.Metric)}
}
