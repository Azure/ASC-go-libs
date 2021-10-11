package instrumentation

import (
	"github.com/sirupsen/logrus"
	"path/filepath"
)

const (
	_metricsDirName = "metrics"
)

// MetricSubmitterFactory is factory for metric submitter
type MetricSubmitterFactory interface {
	// createMetricSubmitter creates MetricSubmitter.
	createMetricSubmitter() MetricSubmitter
}

// MetricSubmitterFactoryImpl implements MetricSubmitterFactory interface
var _ MetricSubmitterFactory = (*MetricSubmitterFactoryImpl)(nil)

// MetricSubmitterFactoryImpl a factory for creating a metric submitter
type MetricSubmitterFactoryImpl struct {
	tracer        *logrus.Entry
	configuration *InstrumentationConfiguration
	metricWriter  MetricWriter
}

// NewMetricSubmitterFactory tracer factory
func NewMetricSubmitterFactory(tracer *logrus.Entry, configuration *InstrumentationConfiguration) (MetricSubmitterFactory, error) {
	metricsPath := filepath.Join(configuration.DirPath, _metricsDirName)
	metricWriter, err := newAggregatedMetricWriter(tracer, configuration.ComponentName, metricsPath)
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
	metricSubmitter := newMetricSubmitter(metricSubmitterFactory.tracer, metricSubmitterFactory.metricWriter, metricSubmitterFactory.configuration.ReleaseTrain,
		metricSubmitterFactory.configuration.ComponentName,
		metricSubmitterFactory.configuration.MdmAccount,
		metricSubmitterFactory.configuration.MdmNamespace,
		metricSubmitterFactory.configuration.GetDefaultDimensions())

	return metricSubmitter
}
