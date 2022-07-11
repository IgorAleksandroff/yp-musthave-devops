package repository

import (
	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics/entity"
)

type rep struct {
	storage map[string]entity.Metric
}

func New() *rep {
	return &rep{storage: make(map[string]entity.Metric)}
}
