package instrumentation

import (
	"strings"
)

// InstrumentationConfiguration for the instrumentation platform
type InstrumentationConfiguration struct {
	ComponentName       string
	AzureResourceID     string
	Region              string
	ClusterDistribution string
	ChartVersion        string
	ImageName           string
	ImageVersion        string
	ReleaseTrain        string
	NodeName            string

	MdmAccount   string
	MdmNamespace string
	// DirPath is the path to the directory that the files will be saved.
	DirPath           string
	PlatformNamespace string
}

// GetDefaultDimensions - Get the default dimensions to be attached to each metric reports
func (configuration *InstrumentationConfiguration) GetDefaultDimensions() []*Dimension {
	return []*Dimension{
		{
			Key:   "ChartVersion",
			Value: configuration.ChartVersion,
		},
		{
			Key:   "ClusterDistribution",
			Value: configuration.ClusterDistribution,
		},
		{
			Key:   "ComponentName",
			Value: configuration.ComponentName,
		},
		{
			Key:   "ImageName",
			Value: configuration.ImageName,
		},
		{
			Key:   "ImageVersion",
			Value: configuration.ImageVersion,
		},
		{
			Key:   "Region",
			Value: strings.ToLower(configuration.Region),
		},
		{
			Key:   "ReleaseTrain",
			Value: strings.ToLower(configuration.ReleaseTrain),
		},
	}
}
