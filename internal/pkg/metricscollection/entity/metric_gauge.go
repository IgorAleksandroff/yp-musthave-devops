package entity

const GaugeTypeMetric = "gauge"
const CounterTypeMetric = "gauge"

type MetricGauge struct {
	TypeMetric string
	Value      float64
}
type MetricCounter struct {
	TypeMetric string
	Value      int64
}
