package instrumentation

// Those are the keys of the environment variables for the tracer and the metric submitter.
const (
	// EnvImageVersionKey Image version key.
	EnvImageVersionKey = "IMAGE_VERSION"
	// EnvResourceIDKey is the key of the id of the resource
	EnvResourceIDKey = "AZURE_RESOURCE_ID"
	// EnvResourceRegionKey is the key of the region of the resource.
	EnvResourceRegionKey = "AZURE_RESOURCE_REGION"
	// EnvClusterDistributionKey is the key of the cluster distribution.
	EnvClusterDistributionKey = "CLUSTER_DISTRIBUTION"
	// EnvChartVersionKey is the key of the chart version.
	EnvChartVersionKey = "CHART_VERSION"
	// EnvNodeNameKey is the key of the node name
	EnvNodeNameKey = "NODE_NAME"
)

const (
	// Unknown is constant that represents unknown string
	Unknown = "Unknown"
	// logFileDir TODO - load it using viper instead of using constant.
	logFileDir = "/var/log/azuredefender/traces/"
)
