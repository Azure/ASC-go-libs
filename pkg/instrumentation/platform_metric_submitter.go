package instrumentation

import "github.com/Azure/Tivan-Libs/pkg/common"

const (
	platformNamespace = "Tivan.Platform"
)

// PlatformMetricSubmitter - interface for sending platform metrics
type PlatformMetricSubmitter interface {
	// SendMetricToPlatform - send metric by name with provided dimensions, to the platform account and namespace
	SendMetricToPlatform(value int, metric Metric)
}

// PlatformMetricSubmitterImpl a metric submitter object - can be use to send metrics easily (to platform account)
type PlatformMetricSubmitterImpl struct {
	underlinedMetricSubmitter MetricSubmitter
	accountName               string
}

// NewPlatformMetricSubmitter creates a new metric submitter that reports metrics to the platform namespace and account
func NewPlatformMetricSubmitter(metricSubmitter MetricSubmitter) PlatformMetricSubmitter {

	region := common.GetEnvVariableOrDefault(EnvResourceRegionKey, Unknown)
	accountName := GetPlatformMdmAccount(region)

	c := &PlatformMetricSubmitterImpl{
		underlinedMetricSubmitter: metricSubmitter,
		accountName:               accountName,
	}

	return c
}

// SendMetric send metric (for platform submitter)
func (platformMetricSubmitter *PlatformMetricSubmitterImpl) SendMetricToPlatform(value int, metric Metric) {
	platformMetricSubmitter.underlinedMetricSubmitter.SendMetricToNamespace(value, metric, platformMetricSubmitter.accountName, platformNamespace)
}
