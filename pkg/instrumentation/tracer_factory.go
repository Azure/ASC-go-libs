package instrumentation

import (
	"github.com/Azure/ASC-go-libs/pkg/common"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	_tracesDirName = "traces"
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
	logFileDir                   string
}

// NewTracerFactory tracer factory
func NewTracerFactory(instrumentationConfiguration *InstrumentationConfiguration) TracerFactory {
	tracesPath := filepath.Join(instrumentationConfiguration.DirPath, _tracesDirName)
	m := &TracerFactoryImpl{
		instrumentationConfiguration: instrumentationConfiguration,
		rollingFileConfiguration:     common.GetDefaultFileConfiguration(),
		logFileDir:                   tracesPath,
	}

	return m
}

// SetRollingFileConfiguration - set rolling file configuration
func (tracerFactory *TracerFactoryImpl) SetRollingFileConfiguration(rollingFileConfiguration *common.RollingFileConfiguration) {
	tracerFactory.rollingFileConfiguration = rollingFileConfiguration
}

// DeleteTracerFile - delete the tracer's log file
func (tracerFactory *TracerFactoryImpl) DeleteTracerFile() error {
	logFilePath := filepath.Join(tracerFactory.logFileDir, tracerFactory.instrumentationConfiguration.ComponentName)
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

	logFilePath := filepath.Join(tracerFactory.logFileDir, tracerFactory.instrumentationConfiguration.ComponentName)

	logFile := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    tracerFactory.rollingFileConfiguration.MaxSize,
		MaxBackups: tracerFactory.rollingFileConfiguration.MaxBackups,
		MaxAge:     tracerFactory.rollingFileConfiguration.MaxAgeInDays,
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetLevel(log.TraceLevel)

	componentVersion := tracerFactory.instrumentationConfiguration.ImageName + ":" + tracerFactory.instrumentationConfiguration.ImageVersion
	entry := log.WithFields(map[string]interface{}{
		"componentVersion":    componentVersion,
		"NodeName":            tracerFactory.instrumentationConfiguration.NodeName,
		"ComponentName":       tracerFactory.instrumentationConfiguration.ComponentName,
		"ReleaseTrain":        tracerFactory.instrumentationConfiguration.ReleaseTrain,
		"AzureResourceID":     tracerFactory.instrumentationConfiguration.AzureResourceID,
		"ChartVersion":        tracerFactory.instrumentationConfiguration.ChartVersion,
		"Region":              tracerFactory.instrumentationConfiguration.Region,
		"ClusterDistribution": tracerFactory.instrumentationConfiguration.ClusterDistribution,
		"type":                tracerType,
	})

	return entry
}
