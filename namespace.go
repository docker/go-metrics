package metrics

import (
	"fmt"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

// NewNamespace returns a namespaces that is responsible for managing a collection of
// metrics for a particual namespace and subsystem
//
// labels allows const labels to be added to all metrics created in this namespace
// and are commonly used for data like application version and git commit
func NewNamespace(name, subsystem string, labels map[string]string) Namespace {
	if labels == nil {
		labels = make(map[string]string)
	}
	return &namespace{
		name:      name,
		subsystem: subsystem,
		labels:    labels,
	}
}

// Namespace is a collection of metrics under a common namespace and subsystem
type Namespace interface {
	// NewCounter returns a new Counter with the provided name and help string along with the keys
	// of any dynamic labels what will be used with the counter
	NewCounter(name, help string, labels []string) Counter
	// NewTimer returns a new Timer with the provided name and help string along with the keys
	// of any dynamic labels what will be used with the timer
	NewTimer(name, help string, labels []string) Timer
	// NewGauge returns a new Gauge with the provided name and help string along with the keys
	// of any dynamic labels what will be used with the gauge
	NewGauge(name, help string, unit Unit, labels []string) Gauge

	getMetrics() []prometheus.Collector
}

type namespace struct {
	name      string
	subsystem string
	labels    map[string]string
	mu        sync.Mutex
	metrics   []prometheus.Collector
}

func (n *namespace) NewCounter(name, help string, labels []string) Counter {
	c := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   n.name,
		Subsystem:   n.subsystem,
		Name:        fmt.Sprintf("%s_%s", name, Total),
		Help:        help,
		ConstLabels: prometheus.Labels(n.labels),
	}, labels)
	n.mu.Lock()
	n.metrics = append(n.metrics, c)
	n.mu.Unlock()
	return &counter{pc: c}
}

func (n *namespace) NewTimer(name, help string, labels []string) Timer {
	t := &timer{
		m: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace:   n.name,
			Subsystem:   n.subsystem,
			Name:        fmt.Sprintf("%s_%s", name, Seconds),
			Help:        help,
			ConstLabels: prometheus.Labels(n.labels),
		}, labels),
	}
	n.mu.Lock()
	n.metrics = append(n.metrics, t)
	n.mu.Unlock()
	return t
}

func (n *namespace) NewGauge(name, help string, unit Unit, labels []string) Gauge {
	g := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   n.name,
		Subsystem:   n.subsystem,
		Name:        fmt.Sprintf("%s_%s", name, unit),
		Help:        help,
		ConstLabels: prometheus.Labels(n.labels),
	}, labels)
	n.mu.Lock()
	n.metrics = append(n.metrics, g)
	n.mu.Unlock()
	return &gauge{pg: g}
}

func (n *namespace) getMetrics() []prometheus.Collector {
	return n.metrics
}
