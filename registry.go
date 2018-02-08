package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	registry = prometheus.NewRegistry()
)
