package runtime_metrics

//go:generate mockery -name Usecase

type Usecase interface {
	UpdateMetrics()
	SendMetrics() error
}
