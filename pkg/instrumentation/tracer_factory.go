package instrumentation

import (
	"github.com/Azure/Tivan-Libs/pkg/common"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type TraceType string

const (
	HEARTBEAT_TRACE_TYPE TraceType = "Heartbeat"
	MESSAGE_TRACE_TYPE   TraceType = "Trace"
)

// TracerFactory - Interface for creating tracer entry objects
type TracerFactory interface {
	// CreateTracer - Method for creating a tracer entry that can be used to send traces
	CreateTracer(tracerType TraceType) *log.Entry
	// DeleteTracerFile - delete the tracer's log file
	DeleteTracerFile() error
}

// TracerFactoryImpl implements TracerFactory interface
var _ TracerFactory = (*TracerFactoryImpl)(nil)

// TracerFactoryImpl a factory for creating a tracer entry
type TracerFactoryImpl struct {
	instrumentationConfiguration *InstrumentationConfiguration
	rollingFileConfiguration     *common.RollingFileConfiguration
}

// NewTracerFactory tracer factory
func NewTracerFactory(instrumentationConfiguration *InstrumentationConfiguration) TracerFactory {
	m := &TracerFactoryImpl{
		instrumentationConfiguration: instrumentationConfiguration,
		rollingFileConfiguration:     common.GetDefaultFileConfiguration(),
	}

	return m
}

// SetRollingFileConfiguration - set rolling file configuration
func (tracerFactory *TracerFactoryImpl) SetRollingFileConfiguration(rollingFileConfiguration *common.RollingFileConfiguration) {
	tracerFactory.rollingFileConfiguration = rollingFileConfiguration
}

// DeleteTracerFile - delete the tracer's log file
func (tracerFactory *TracerFactoryImpl) DeleteTracerFile() error {
	logFilePath := logFileDir + tracerFactory.instrumentationConfiguration.componentName
	return os.Remove(logFilePath)
}

// CreateTracer - Method for creating a tracer entry that can be used to send traces
func (tracerFactory *TracerFactoryImpl) CreateTracer(tracerType TraceType) *log.Entry {
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyMsg:   "message",
			log.FieldKeyLevel: "traceLevel",
			log.FieldKeyTime:  "envTime",
		},
	})

	logFilePath := logFileDir + tracerFactory.instrumentationConfiguration.componentName

	logFile := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    tracerFactory.rollingFileConfiguration.MaxSize,
		MaxBackups: tracerFactory.rollingFileConfiguration.MaxBackups,
		MaxAge:     tracerFactory.rollingFileConfiguration.MaxAgeInDays,
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetLevel(log.TraceLevel)

	componentVersion := tracerFactory.instrumentationConfiguration.imageName + ":" + tracerFactory.instrumentationConfiguration.imageVersion
	entry := log.WithFields(map[string]interface{}{
		"componentVersion":    componentVersion,
		"nodeName":            tracerFactory.instrumentationConfiguration.nodeName,
		"componentName":       tracerFactory.instrumentationConfiguration.componentName,
		"releaseTrain":        tracerFactory.instrumentationConfiguration.releaseTrain,
		"azureResourceID":     tracerFactory.instrumentationConfiguration.azureResourceID,
		"chartVersion":        tracerFactory.instrumentationConfiguration.chartVersion,
		"region":              tracerFactory.instrumentationConfiguration.region,
		"clusterDistribution": tracerFactory.instrumentationConfiguration.clusterDistribution,
		"type":                tracerType,
	})

	return entry
}
