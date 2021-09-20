package instrumentation

import (
	log "github.com/sirupsen/logrus"
)

// MetricSubmitter - interface for sending metrics
type MetricSubmitter interface {
	// SendMetric - send metric by name with provided dimensions
	SendMetric(value int, metric Metric)

	SendMetricToNamespace(value int, metric Metric, accountName, namespaceName string)
}

// MetricSubmitterImpl a metric submitter object - can be use to send metrics easily
type MetricSubmitterImpl struct {
	tracer            *log.Entry
	releaseTrain      string
	accountName       string
	namespaceName     string
	componentName     string
	defaultDimensions []Dimension
	metricWriter      MetricWriter
}

// NewMetricSubmitter creates a new metric submitter
func newMetricSubmitter(tracer *log.Entry, metricWriter MetricWriter, releaseTrain, componentName, accountName, namespaceName string, defaultDimensions []Dimension) MetricSubmitter {
	c := &MetricSubmitterImpl{
		tracer:            tracer,
		releaseTrain:      releaseTrain,
		accountName:       accountName,
		namespaceName:     namespaceName,
		componentName:     componentName,
		defaultDimensions: defaultDimensions,
		metricWriter:      metricWriter,
	}

	return c
}

// SendMetric send metric
func (metricSubmitter *MetricSubmitterImpl) SendMetric(value int, metric Metric) {
	metricSubmitter.SendMetricToNamespace(value, metric, metricSubmitter.accountName, metricSubmitter.namespaceName)
}

func (metricSubmitter *MetricSubmitterImpl) SendMetricToNamespace(value int, metric Metric, accountName, namespaceName string) {
	metricDimensions := metricSubmitter.getDimensionsToSend(metric.MetricDimension())
	metricToSend := newRawMetric(metricSubmitter.releaseTrain, metricSubmitter.componentName,
		accountName, namespaceName, metric.MetricName(), metricDimensions, uint32(value))

	metricSubmitter.metricWriter.Write(metricToSend)
}

func (metricSubmitter MetricSubmitterImpl) getDimensionsToSend(dimensions []Dimension) []Dimension {
	mergedDimensions := make([]Dimension, len(metricSubmitter.defaultDimensions))
	copy(mergedDimensions, metricSubmitter.defaultDimensions)

	for _, dimension := range dimensions {
		found := false
		for i, existingDimension := range mergedDimensions {
			if existingDimension.Key == dimension.Key {
				mergedDimensions[i].Value = dimension.Value
				found = true
				break
			}
		}
		if !found {
			mergedDimensions = append(mergedDimensions, dimension)
		}
	}

	return mergedDimensions
}
