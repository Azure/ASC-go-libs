package instrumentation

const (
	// Metric name for event read
	eventReadMetricName = "EventRead"

	// Metric name for event parsed
	eventParsedMetricName = "EventParsed"

	// Metric name for event enriched
	eventEnrichedMetricName = "EventEnriched"

	// Metric name for event collected
	eventCollectedMetricName = "EventCollected"

	// Dimension name
	eventTypeDimensionKey = "EventType"
)

// EventBasedMetric implementation of Metric, for event based metric
type EventBasedMetric struct {
	eventName string
	eventType string
}

// MetricName - metric name
func (metric *EventBasedMetric) MetricName() string {
	return metric.eventName
}

// MetricDimension - metric dimensions
func (metric *EventBasedMetric) MetricDimension() []Dimension {
	return []Dimension{
		{
			Key:   eventTypeDimensionKey,
			Value: metric.eventType,
		},
	}
}

// NewEventReadMetric Ctor for EventReadMetric
func NewEventReadMetric(eventType string) *EventBasedMetric {
	return &EventBasedMetric{
		eventName: eventReadMetricName,
		eventType: eventType,
	}
}

// NewEventParsedMetric Ctor for EventParsedMetric
func NewEventParsedMetric(eventType string) *EventBasedMetric {
	return &EventBasedMetric{
		eventName: eventParsedMetricName,
		eventType: eventType,
	}
}

// NewEventEnrichedMetric Ctor for EventEnrichedMetric
func NewEventEnrichedMetric(eventType string) *EventBasedMetric {
	return &EventBasedMetric{
		eventName: eventEnrichedMetricName,
		eventType: eventType,
	}
}

// newEventCollectedMetric Ctor for EventCollectedMetric
func NewEventCollectedMetric(eventType string) *EventBasedMetric {
	return &EventBasedMetric{
		eventName: eventCollectedMetricName,
		eventType: eventType,
	}
}
