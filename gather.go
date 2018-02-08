package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

// Gatherer returns the metric gatherer
func Gatherer() prometheus.Gatherer {
	return registry
}

// Gather calls the gatherer to return all the metric families of
// every metric in the registry.
func Gather() ([]*dto.MetricFamily, error) {
	return Gatherer().Gather()
}
