package metricscollection

//go:generate mockery --name Usecase

type Usecase interface {
	SaveGaugeMetric(name string, value float64)
	SaveCounterMetric(name string, value int64)
}
