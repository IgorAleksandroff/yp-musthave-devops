package chat_messages

//go:generate mockery -name Repository

type Repository interface {
	SaveGaugeMetric(name string, value float64) error
}
