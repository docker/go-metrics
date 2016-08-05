package metrics

import "github.com/prometheus/client_golang/prometheus"

// Gauge is a metric that allows incrementing and decrementing a value
type Gauge interface {
	// Add adds the provided value to the gauge's current value
	Add(float64, map[string]string)
	// Set replaces the gauge's current value with the provided value
	Set(float64, map[string]string)
}

type gauge struct {
	pg *prometheus.GaugeVec
}

func (g *gauge) Add(v float64, labels map[string]string) {
	g.pg.With(prometheus.Labels(labels)).Add(v)
}

func (g *gauge) Set(v float64, labels map[string]string) {
	g.pg.With(prometheus.Labels(labels)).Set(v)
}
