package instrumentation

// PlatformMetricSubmitter - interface for sending platform metrics
type PlatformMetricSubmitter interface {
	// SendMetricToPlatform - send metric by name with provided dimensions, to the platform account and namespace
	SendMetricToPlatform(value int, metric Metric)
}

// PlatformMetricSubmitterImpl implements PlatformMetricSubmitter interface
var _ PlatformMetricSubmitter = (*PlatformMetricSubmitterImpl)(nil)

// PlatformMetricSubmitterImpl a metric submitter object - can be used to send metrics easily (to platform account)
type PlatformMetricSubmitterImpl struct {
	underlinedMetricSubmitter MetricSubmitter
	accountName               string
	namespace                 string
}

// NewPlatformMetricSubmitter creates a new metric submitter that reports metrics to the platform namespace and account
func NewPlatformMetricSubmitter(metricSubmitter MetricSubmitter, mdmAccount string, namespace string) PlatformMetricSubmitter {

	c := &PlatformMetricSubmitterImpl{
		underlinedMetricSubmitter: metricSubmitter,
		accountName:               mdmAccount,
		namespace:                 namespace,
	}

	return c
}

// SendMetricToPlatform send metric (for platform submitter)
func (platformMetricSubmitter *PlatformMetricSubmitterImpl) SendMetricToPlatform(value int, metric Metric) {
	platformMetricSubmitter.underlinedMetricSubmitter.SendMetricToNamespace(value, metric, platformMetricSubmitter.accountName, platformMetricSubmitter.namespace)
}
