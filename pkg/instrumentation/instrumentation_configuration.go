package instrumentation

import (
	"github.com/Azure/Tivan-Libs/pkg/common"
	"strings"
)

const (
	_numberOfImageVersionComponents = 4
	_defaultReleaseTrain            = "dev"
	_mcr                            = "mcr"
)

var (
	_availableReleaseTrain = []string{"dev", "stage", "stable"}
)

// InstrumentationConfiguration for the instrumentation platform
type InstrumentationConfiguration struct {
	componentName       string
	azureResourceID     string
	region              string
	clusterDistribution string
	chartVersion        string
	imageName           string
	imageVersion        string
	releaseTrain        string
	nodeName            string

	mdmAccount   string
	mdmNamespace string
}

// NewInstrumentationConfiguration - Ctor to create a new instrumentation configuration
func NewInstrumentationConfiguration(componentName, azureResourceID, region, clusterDistribution, chartVersion, imageName, imageVersion, mdmAccount, mdmNamespace, releaseTrain, nodeName string) *InstrumentationConfiguration {
	return &InstrumentationConfiguration{
		componentName:       componentName,
		azureResourceID:     azureResourceID,
		region:              region,
		clusterDistribution: clusterDistribution,
		chartVersion:        chartVersion,
		imageName:           imageName,
		imageVersion:        imageVersion,
		mdmAccount:          mdmAccount,
		mdmNamespace:        mdmNamespace,
		releaseTrain:        releaseTrain,
		nodeName:            nodeName,
	}
}

// NewInstrumentationConfigurationFromEnv - Ctor to create a new instrumentation configuration from environment variables
func NewInstrumentationConfigurationFromEnv(componentName, mdmNamespace string) *InstrumentationConfiguration {
	azureResourceID := common.GetEnvVariableOrDefault(EnvResourceIDKey, Unknown)
	region := common.GetEnvVariableOrDefault(EnvResourceRegionKey, Unknown)
	clusterDistribution := common.GetEnvVariableOrDefault(EnvClusterDistributionKey, Unknown)
	chartVersion := common.GetEnvVariableOrDefault(EnvChartVersionKey, Unknown)
	nodeName := common.GetEnvVariableOrDefault(EnvNodeNameKey, Unknown)
	imageName, imageVersion := splitImageAndVersion(common.GetEnvVariableOrDefault(EnvImageVersionKey, Unknown))
	releaseTrain := GetReleaseTrain(region)

	// TODO: support overriding of the mdmAccount and using account which is different from the platform
	mdmAccount := GetPlatformMdmAccount(region)

	return NewInstrumentationConfiguration(componentName, azureResourceID, region, clusterDistribution, chartVersion,
		imageName, imageVersion, mdmAccount, mdmNamespace, releaseTrain, nodeName)
}

// GetDefaultDimensions - Get the default dimensions to be attached to each metric reports
func (configuration *InstrumentationConfiguration) GetDefaultDimensions() []*Dimension {
	return []*Dimension{
		{
			Key:   "ChartVersion",
			Value: configuration.chartVersion,
		},
		{
			Key:   "ClusterDistribution",
			Value: configuration.clusterDistribution,
		},
		{
			Key:   "ComponentName",
			Value: configuration.componentName,
		},
		{
			Key:   "ImageName",
			Value: configuration.imageName,
		},
		{
			Key:   "ImageVersion",
			Value: configuration.imageVersion,
		},
		{
			Key:   "Region",
			Value: strings.ToLower(configuration.region),
		},
		{
			Key:   "ReleaseTrain",
			Value: strings.ToLower(configuration.releaseTrain),
		},
	}
}

func splitImageAndVersion(imageAndVersion string) (imageName, imageVersion) {
	imageAndVersionArray := strings.Split(imageAndVersion, ":")
	if len(imageAndVersionArray) == 2 {
		return imageAndVersionArray[0], imageAndVersionArray[1]
	} else {
		return imageAndVersion, imageAndVersion
	}
}
