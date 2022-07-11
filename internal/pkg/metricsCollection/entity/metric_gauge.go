package entity

const TypeMetric = "gauge"

type MetricGauge struct {
	TypeMetric string
	Value      float64
}
