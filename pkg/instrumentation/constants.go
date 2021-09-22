package instrumentation

// Those are the default value for the tracer and the metric submitter.
const (
	EnvImageVersionKey        = "IMAGE_VERSION"
	EnvResourceIDKey          = "AZURE_RESOURCE_ID"
	EnvResourceRegionKey      = "AZURE_RESOURCE_REGION"
	EnvClusterDistributionKey = "CLUSTER_DISTRIBUTION"
	EnvChartVersionKey        = "CHART_VERSION"
	EnvNodeNameKey            = "NODE_NAME"
	Unknown                   = "Unknown"
	logFileDir                = "/var/log/azuredefender/traces/"
)
