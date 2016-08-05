package metrics

// Unit represents the type or precision of a metric that is appended to
// the metrics fully qualified name
type Unit string

const (
	// Nanoseconds specifies that the metric's values are in nanoseconds
	Nanoseconds Unit = "ns"
	// Milliseconds specifies that the metric's values are in milliseconds
	Milliseconds Unit = "ms"
	// Seconds specifies that the metric's values are in seconds
	Seconds Unit = "s"
	// Count specifies that the metric's values are a count
	Count Unit = "count"
)
