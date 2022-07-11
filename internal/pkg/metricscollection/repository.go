package metricscollection

//go:generate mockery -name Repository

type Repository interface {
	SaveGaugeMetric(name string, value float64)
}
