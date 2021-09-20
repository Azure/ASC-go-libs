package common

import (
	"encoding/json"

	"github.com/tdewolff/minify"
	jsonMinifer "github.com/tdewolff/minify/json"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	defaultRollingFileConfiguration = NewRollingFileConfiguration(1, 1, 1)
	jsonMimeType                    = "application/json"
)

// GetDefaultFileConfiguration - return default configuration
func GetDefaultFileConfiguration() *RollingFileConfiguration {
	return defaultRollingFileConfiguration
}

// FileWriter - an interface for writing to files
type FileWriter interface {
	// Write - write the interface to the file
	Write(data interface{}) error
	Close()
}

// RollingFileConfiguration - configuration for rolling file writer
type RollingFileConfiguration struct {
	MaxSize      int
	MaxBackups   int
	MaxAgeInDays int
}

// NewRollingFileConfiguration - Ctor to create a new Rolling file configuration
func NewRollingFileConfiguration(maxSize, maxBackups, maxAgeInDays int) *RollingFileConfiguration {
	return &RollingFileConfiguration{
		MaxSize:      maxSize,
		MaxBackups:   maxBackups,
		MaxAgeInDays: maxAgeInDays,
	}
}

// RollingFileWriter - helper object can be use to write objects to rolling files
// RollingFileWriter implements FileWriter interface
type RollingFileWriter struct {
	filePath string
	minifier *minify.M
	file     *lumberjack.Logger
}

// NewRollingFileWriter - Ctor to create a new Rolling file writer
// maxSize - megabytes after which new file is created
// maxBackups - number of backups
// maxAge - max time in days before removing backup
func NewRollingFileWriter(filePath string) *RollingFileWriter {

	var minifier = minify.New()
	minifier.AddFunc(jsonMimeType, jsonMinifer.Minify)

	m := &RollingFileWriter{
		filePath: filePath,
		minifier: minifier,
		file:     getRollingFile(filePath, defaultRollingFileConfiguration),
	}

	return m
}

// SetRollingFileConfiguration - set rolling file configuration
func (rollingFileWriter *RollingFileWriter) SetRollingFileConfiguration(configuration *RollingFileConfiguration) {
	file := getRollingFile(rollingFileWriter.filePath, configuration)

	rollingFileWriter.file = file
}

// Write data to file
func (rollingFileWriter *RollingFileWriter) Write(data interface{}) error {
	var out []byte
	var err error

	switch val := data.(type) {
	case string:
		out = []byte(val)
	default:
		notMinifiedOutput, _ := json.Marshal(data)
		out, err = rollingFileWriter.minifier.Bytes(jsonMimeType, notMinifiedOutput)

		if err != nil {
			return err
		}
	}

	_, err = rollingFileWriter.file.Write([]byte(string(out) + "\n"))

	return err
}

// Close the writer
func (rollingFileWriter *RollingFileWriter) Close() {
	rollingFileWriter.file.Close()
}

func getRollingFile(filePath string, configuration *RollingFileConfiguration) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    configuration.MaxSize,
		MaxBackups: configuration.MaxBackups,
		MaxAge:     configuration.MaxAgeInDays,
	}
}
