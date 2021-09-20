package instrumentation

// Dimension A dimension in the metric
type Dimension struct {
	Key   string
	Value string
}

// Metric interface for getting the metric name and metric dimensions
type Metric interface {
	// GetMetricName - getter for the metric name
	MetricName() string
	// GetMetricDimension - getter for the metric dimensions
	MetricDimension() []Dimension
}
