package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Timer is a metric that allows collecting the duration of an action in seconds
type Timer interface {
	// UpdateSince will add the duration from the provided starting time to the
	// timer's summary with the precisions that was used in creation of the timer
	UpdateSince(time.Time, map[string]string)
}

type timer struct {
	m *prometheus.HistogramVec
}

func (t *timer) UpdateSince(since time.Time, labels map[string]string) {
	t.m.With(prometheus.Labels(labels)).Observe(time.Now().Sub(since).Seconds())
}

func (t *timer) Describe(c chan<- *prometheus.Desc) {
	t.m.Describe(c)
}

func (t *timer) Collect(c chan<- prometheus.Metric) {
	t.m.Collect(c)
}
