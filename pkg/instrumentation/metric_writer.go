package instrumentation

import (
	"github.com/Azure/ASC-go-libs/pkg/common"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	_reportingInterval = 1 * time.Minute
)

// MetricWriter - interface for sending metrics
type MetricWriter interface {
	// Write - send metric by name with provided dimensions
	Write(metric *rawMetric)
}

// MetricWriterImpl implements MetricWriter interface
var _ MetricWriter = (*MetricWriterImpl)(nil)

// MetricWriterImpl a metric sender object - can be used to send metrics easily
type MetricWriterImpl struct {
	tracer     *log.Entry
	fileWriter common.FileWriter
	lock       sync.RWMutex
	values     map[uint32]*rawMetric
}

// newAggregatedMetricWriter creates a new metric writer aggregator
func newAggregatedMetricWriter(tracer *log.Entry, componentName string, dirPath string) (MetricWriter, error) {
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, err
	}

	filePath := filepath.Join(dirPath, componentName)
	_, err := os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
	if err != nil {
		tracer.Error(err)
		return nil, err
	}

	c := &MetricWriterImpl{
		tracer:     tracer,
		fileWriter: common.NewRollingFileWriter(filePath),
		lock:       sync.RWMutex{},
		values:     make(map[uint32]*rawMetric),
	}

	go c.startReporting()

	return c, nil
}

// Write - writes the new metric or updating the value of the existing cached one
func (metricWriter *MetricWriterImpl) Write(metric *rawMetric) {

	metricWriter.lock.RLock()

	hash := metric.GetHashExcludingValue()

	if _, ok := metricWriter.values[hash]; !ok {
		metricWriter.lock.RUnlock()
		metricWriter.lock.Lock()
		defer metricWriter.lock.Unlock()
		metricWriter.values[hash] = metric
	} else {
		defer metricWriter.lock.RUnlock()

		atomic.AddUint32(&metricWriter.values[hash].Value, metric.Value)
	}
}

func (metricWriter *MetricWriterImpl) startReporting() {
	ticker := time.NewTicker(_reportingInterval)
	for range ticker.C {
		metricWriter.report()
	}
}

func (metricWriter *MetricWriterImpl) report() {
	metricWriter.lock.Lock()
	defer metricWriter.lock.Unlock()
	for _, value := range metricWriter.values {
		err := metricWriter.fileWriter.Write(value)

		if err != nil {
			metricWriter.tracer.Error("Failed to send metric", err)
		}
	}
	metricWriter.values = make(map[uint32]*rawMetric)
}
