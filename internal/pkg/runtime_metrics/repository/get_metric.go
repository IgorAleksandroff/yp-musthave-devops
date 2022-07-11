package repository

import (
	"fmt"

	"github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics/entity"
)

func (r rep) GetMetric(name string) (m entity.Metric, err error) {
	if metric, ok := r.storage[name]; ok {
		return metric, nil
	}

	return m, fmt.Errorf("can not get a metric: %s", name)
}
