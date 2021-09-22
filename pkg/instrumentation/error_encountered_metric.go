package instrumentation

const (
	// ErrorEncounteredMetricName is the metric name
	ErrorEncounteredMetricName = "ErrorEncountered"

	// Dimensions names
	errorTypeDimensionKey = "ErrorType"
	contextDimensionKey   = "Context"
)

// ErrorEncounteredMetric implementation of Metric, for error encountered metric
type ErrorEncounteredMetric struct {
	// errorType is the type of the error that was encountered
	errorType string
	// context of the metric.
	context string
}

// NewErrorEncounteredMetric Cto'r for ErrorEncounteredMetric
func NewErrorEncounteredMetric(errorType, context string) *ErrorEncounteredMetric {
	return &ErrorEncounteredMetric{
		errorType: errorType,
		context:   context,
	}
}

func (metric *ErrorEncounteredMetric) MetricName() string {
	return ErrorEncounteredMetricName
}

func (metric *ErrorEncounteredMetric) MetricDimension() []Dimension {
	return []Dimension{
		{
			Key:   errorTypeDimensionKey,
			Value: metric.errorType,
		},
		{
			Key:   contextDimensionKey,
			Value: metric.context,
		},
	}
}
