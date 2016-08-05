package metrics

// Suffix represents the type or precision of a metric that is appended to
// the metrics fully qualified name
type Suffix string

const (
	// Nanoseconds specifies that the metric's values are in nanoseconds
	Nanoseconds Suffix = "ns"
	// Milliseconds specifies that the metric's values are in milliseconds
	Milliseconds Suffix = "ms"
	// Seconds specifies that the metric's values are in seconds
	Seconds Suffix = "s"
	// Count specifies that the metric's values are a count
	Count Suffix = "count"
)
