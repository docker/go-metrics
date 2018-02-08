package metrics

// Register adds all the metrics in the provided namespace to the global
// metrics registry
func Register(n *Namespace) {
	registry.MustRegister(n)
}

// Deregister removes all the metrics in the provided namespace from the
// global metrics registry
func Deregister(n *Namespace) {
	registry.Unregister(n)
}
