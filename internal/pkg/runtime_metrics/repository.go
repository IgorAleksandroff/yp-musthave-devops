package runtime_metrics

import "github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics/entity"

//go:generate mockery -name Repository

type Repository interface {
	SaveMetric(name string, value entity.Getter)
	GetMetric(name string) (m entity.Metric, err error)
	GetMetricsName() []string
}
