package instrumentation

// DimensionlessMetric implementation of Metric, for metric without dimensions
type DimensionlessMetric struct {
	metricName string
}

// NewDimensionlessMetric Cto'r for DimensionlessMetric
func NewDimensionlessMetric(metricName string) *DimensionlessMetric {
	return &DimensionlessMetric{
		metricName: metricName,
	}
}

func (metric *DimensionlessMetric) MetricName() string {
	return metric.metricName
}

func (metric *DimensionlessMetric) MetricDimension() []Dimension {
	return []Dimension{}
}
