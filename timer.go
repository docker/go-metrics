package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// StartTimer begins a timer observation at the callsite. When the target
// operation is completed, the caller should call the return done func().
func StartTimer(timer Timer, labels map[string]string) (done func()) {
	start := time.Now()
	return func() {
		timer.Update(time.Since(start), labels)
	}
}

// Timer is a metric that allows collecting the duration of an action in seconds
type Timer interface {
	// Update records an observation, duration, and converts to the target
	// units.
	Update(duration time.Duration, labels map[string]string)

	// UpdateSince will add the duration from the provided starting time to the
	// timer's summary with the precisions that was used in creation of the timer
	UpdateSince(time.Time, map[string]string)
}

type timer struct {
	m *prometheus.HistogramVec
}

func (t *timer) Update(duration time.Duration, labels map[string]string) {
	t.m.With(prometheus.Labels(labels)).Observe(duration.Seconds())
}

func (t *timer) UpdateSince(since time.Time, labels map[string]string) {
	t.m.With(prometheus.Labels(labels)).Observe(time.Since(since).Seconds())
}

func (t *timer) Describe(c chan<- *prometheus.Desc) {
	t.m.Describe(c)
}

func (t *timer) Collect(c chan<- prometheus.Metric) {
	t.m.Collect(c)
}
