package runtimeMetrics

//go:generate mockery -name Usecase

type Usecase interface {
	UpdateMetrics()
	SendMetrics() error
}
