package instrumentation

const (
	_canaryRegion       = "eastus2euap"
	_prodMdmAccount     = "RomeDetection"
	_stageMdmAccount    = "RomeDetectionStage"
	_stageReleaseTrain  = "stage"
	_stableReleaseTrain = "stable"
)

// InitializeFromEnv initialize instrumentation using environment variables and returns InstrumentationInitializationResult
func InitializeFromEnv(componentName, mdmNamespace string) *InstrumentationInitializationResult {
	configuration := NewInstrumentationConfigurationFromEnv(componentName, mdmNamespace)

	instrumentationInitializer := NewInstrumentationInitializer(configuration)
	instrumentationInitializationResults, err := instrumentationInitializer.Initialize()

	if err != nil {
		instrumentationInitializationResults.Tracer.Panic("error encountered during tracer initialization", err)
	}

	return instrumentationInitializationResults
}

// GetPlatformMdmAccount returns platform mdm account.
// if region == "eastus2euap" (_canaryRegion) returns "RomeDetectionStage" (_stageMdmAccount),
// else returns  "RomeDetection"( _prodMdmAccount)
func GetPlatformMdmAccount(region string) string {
	if region == _canaryRegion {
		return _stageMdmAccount
	} else {
		return _prodMdmAccount
	}
}

// GetReleaseTrain returns release train
func GetReleaseTrain(region string) string {
	if region == _canaryRegion {
		return _stageReleaseTrain
	} else {
		return _stableReleaseTrain
	}
}
