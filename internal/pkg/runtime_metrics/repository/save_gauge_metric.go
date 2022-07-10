package repository

import "github.com/IgorAleksandroff/yp-musthave-devops/internal/pkg/runtime_metrics/entity"

func (r rep) SaveGaugeMetric(name string, value float64) {
	r.storage[name] = entity.MetricGauge{
		TypeMetric: entity.TypeMetric,
		Value:      value,
	}
}
