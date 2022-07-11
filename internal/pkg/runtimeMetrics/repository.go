package runtimeMetrics

import "github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtimeMetrics/entity"

//go:generate mockery -name Repository

type Repository interface {
	SaveMetric(name string, value entity.Getter)
	GetMetric(name string) (m entity.Metric, err error)
	GetMetricsName() []string
}
