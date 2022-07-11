package metricsCollection

//go:generate mockery -name Usecase

type Usecase interface {
	SaveGaugeMetric(name string, value float64)
}
