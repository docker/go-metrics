package metrics

import "github.com/prometheus/client_golang/prometheus"

// Counter is a metrics that can only increment its current count
type Counter interface {
	// Inc increments the counter's value by 1
	Inc(map[string]string)
	// Add will add the provided value to the counter's current value
	Add(float64, map[string]string)
}

type counter struct {
	pc *prometheus.CounterVec
}

func (c *counter) Inc(labels map[string]string) {
	c.Add(1, labels)
}

func (c *counter) Add(v float64, labels map[string]string) {
	c.pc.With(prometheus.Labels(labels)).Add(v)
}
