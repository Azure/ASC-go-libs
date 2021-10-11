package instrumentation

const (
	// _eventReadMetricName is Metric name for event read
	_eventReadMetricName = "EventRead"

	// _eventParsedMetricName is metric name for event parsed
	_eventParsedMetricName = "EventParsed"

	// _eventEnrichedMetricName is Metric name for event enriched
	_eventEnrichedMetricName = "EventEnriched"

	// _eventCollectedMetricName is Metric name for event collected
	_eventCollectedMetricName = "EventCollected"

	// _eventTypeDimensionKey is the Dimension name
	_eventTypeDimensionKey = "EventType"
)

// EventBasedMetric implements Metric interface
var _ Metric = (*EventBasedMetric)(nil)

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
func (metric *EventBasedMetric) MetricDimension() []*Dimension {
	return []*Dimension{
		{
			Key:   _eventTypeDimensionKey,
			Value: metric.eventType,
		},
	}
}

// NewEventReadMetric Ctor for EventReadMetric
func NewEventReadMetric(eventType string) *EventBasedMetric {
	return &EventBasedMetric{
		eventName: _eventReadMetricName,
		eventType: eventType,
	}
}

// NewEventParsedMetric Ctor for EventParsedMetric
func NewEventParsedMetric(eventType string) *EventBasedMetric {
	return &EventBasedMetric{
		eventName: _eventParsedMetricName,
		eventType: eventType,
	}
}

// NewEventEnrichedMetric Ctor for EventEnrichedMetric
func NewEventEnrichedMetric(eventType string) *EventBasedMetric {
	return &EventBasedMetric{
		eventName: _eventEnrichedMetricName,
		eventType: eventType,
	}
}

// NewEventCollectedMetric Ctor for EventCollectedMetric
func NewEventCollectedMetric(eventType string) *EventBasedMetric {
	return &EventBasedMetric{
		eventName: _eventCollectedMetricName,
		eventType: eventType,
	}
}
