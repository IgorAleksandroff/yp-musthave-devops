package repository

import "github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtimeMetrics/entity"

func (r rep) SaveMetric(name string, value entity.Getter) {
	r.storage[name] = entity.Metric{
		Value: value,
	}
}
