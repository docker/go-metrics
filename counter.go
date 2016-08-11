package metrics

import "github.com/prometheus/client_golang/prometheus"

// Counter is a metrics that can only increment its current count
type Counter interface {
	// Inc the counter by 1 or by the values provided in vs.
	// Sum(vs) must be greater than zero.
	Inc(labels map[string]string, vs ...float64)
}

type counter struct {
	pc *prometheus.CounterVec
}

func (c *counter) Inc(labels map[string]string, vs ...float64) {
	var sum float64
	if len(vs) == 0 {
		sum = 1
	}

	for _, v := range vs {
		sum += v
	}

	c.pc.With(prometheus.Labels(labels)).Add(sum)
}
