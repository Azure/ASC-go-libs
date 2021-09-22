package instrumentation

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	_heartbeatInterval = time.Minute * 5
)

// InstrumentationInitializer helper object to initialize the instrumentation platform
type InstrumentationInitializer struct {
	configuration *InstrumentationConfiguration
}

// InstrumentationInitializationResult - aggregate type for all initialization result objects
type InstrumentationInitializationResult struct {
	MetricSubmitter MetricSubmitter
	Tracer          *log.Entry
}

// NewInstrumentationInitializer - Ctor for creating the instrumentation initializer
func NewInstrumentationInitializer(configuration *InstrumentationConfiguration) *InstrumentationInitializer {
	return &InstrumentationInitializer{
		configuration: configuration,
	}
}

// Initialize - initialize the instrumentation framework
func (initializer *InstrumentationInitializer) Initialize() (*InstrumentationInitializationResult, error) {
	tracerFactory := NewTracerFactory(initializer.configuration)

	tracer := tracerFactory.CreateTracer(MESSAGE_TRACE_TYPE)

	metricSubmitterFactory, err := NewMetricSubmitterFactory(tracer, initializer.configuration)
	if err != nil {
		return nil, err
	}

	metricSubmitter := metricSubmitterFactory.createMetricSubmitter()

	heartbeatTracer := tracerFactory.CreateTracer(HEARTBEAT_TRACE_TYPE)

	if err != nil {
		return nil, err
	}

	platformMetricSubmitter := NewPlatformMetricSubmitter(metricSubmitter)
	heartbeatSender := newHeartbeatSender(heartbeatTracer, platformMetricSubmitter, _heartbeatInterval)
	heartbeatSender.start()

	return &InstrumentationInitializationResult{
		MetricSubmitter: metricSubmitter,
		Tracer:          tracer,
	}, nil
}
