package usecase

func (u usecase) SaveGaugeMetric(name string, value float64) error {
	return u.repository.SaveGaugeMetric(name, value)
}
