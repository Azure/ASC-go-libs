package instrumentation

// Metric interface for getting the metric name and metric dimensions
type Metric interface {
	// MetricName - getter for the metric name
	MetricName() string
	// MetricDimension - getter for the metric dimensions
	MetricDimension() []*Dimension
}

// Dimension A dimension in the metric
type Dimension struct {
	// Key of the dimension.
	Key string
	// Value of the dimension.
	Value string
}
