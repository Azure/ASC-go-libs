package instrumentation

import (
	"github.com/sirupsen/logrus"
)

type imageName = string
type imageVersion = string

type MetricSubmitterFactory interface {
	createMetricSubmitter() MetricSubmitter
}

// MetricSubmitterFactoryImpl a factory for creating a metric submitter
type MetricSubmitterFactoryImpl struct {
	tracer        *logrus.Entry
	configuration *InstrumentationConfiguration
	metricWriter  MetricWriter
}

// NewMetricSubmitterFactory tracer factory
func NewMetricSubmitterFactory(tracer *logrus.Entry, configuration *InstrumentationConfiguration) (MetricSubmitterFactory, error) {
	metricWriter, err := newAggregatedMetricWriter(tracer, configuration.componentName)
	if err != nil {
		tracer.Fatal(err)
		return nil, err
	}

	m := &MetricSubmitterFactoryImpl{
		tracer:        tracer,
		configuration: configuration,
		metricWriter:  metricWriter,
	}

	return m, nil
}

// createMetricSubmitter - Create general metric submitter
func (metricSubmitterFactory *MetricSubmitterFactoryImpl) createMetricSubmitter() MetricSubmitter {
	metricSubmitter := newMetricSubmitter(metricSubmitterFactory.tracer, metricSubmitterFactory.metricWriter, metricSubmitterFactory.configuration.releaseTrain,
		metricSubmitterFactory.configuration.componentName,
		metricSubmitterFactory.configuration.mdmAccount,
		metricSubmitterFactory.configuration.mdmNamespace,
		metricSubmitterFactory.configuration.GetDefaultDimensions())

	return metricSubmitter
}
