package repository

func (r rep) SaveGaugeMetric(name string, value float64) error {
	rep.metricSorage[string] = value
	return nil
}
