package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Timer is a metric that allows collecting the duration of some action
type Timer interface {
	// UpdateSince will add the duration from the provided starting time to the
	// timer's summary with the precisions that was used in creation of the timer
	UpdateSince(time.Time, map[string]string)
}

type timer struct {
	m    *prometheus.SummaryVec
	unit Unit
}

func (t *timer) UpdateSince(since time.Time, labels map[string]string) {
	var (
		v float64
		d = time.Now().Sub(since)
	)
	switch t.unit {
	case Nanoseconds:
		v = float64(d.Nanoseconds())
	case Milliseconds:
		v = float64(d.Nanoseconds() / 1e6)
	case Seconds:
		v = float64(d.Seconds())
	}
	t.m.With(prometheus.Labels(labels)).Observe(v)
}

func (t *timer) Describe(c chan<- *prometheus.Desc) {
	t.m.Describe(c)
}

func (t *timer) Collect(c chan<- prometheus.Metric) {
	t.m.Collect(c)
}
