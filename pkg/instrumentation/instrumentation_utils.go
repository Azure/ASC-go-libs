package instrumentation

const (
	_canaryRegion       = "eastus2euap"
	_prodMdmAccount     = "RomeDetection"
	_stageMdmAccount    = "RomeDetectionStage"
	_stageReleaseTrain  = "stage"
	_stableReleaseTrain = "stable"
)

// InitializeFromEnv TODO
func InitializeFromEnv(componentName, mdmNamespace string) *InstrumentationInitializationResult {
	configuration := NewInstrumentationConfigurationFromEnv(componentName, mdmNamespace)

	instrumentationInitializer := NewInstrumentationInitializer(configuration)
	instrumentationInitializationResults, err := instrumentationInitializer.Initialize()

	if err != nil {
		instrumentationInitializationResults.Tracer.Panic("error encountered during tracer initialization", err)
	}

	return instrumentationInitializationResults
}

// GetPlatformMdmAccount TODO
func GetPlatformMdmAccount(region string) string {
	if region == _canaryRegion {
		return _stageMdmAccount
	} else {
		return _prodMdmAccount
	}
}

// GetReleaseTrain TODO
func GetReleaseTrain(region string) string {
	if region == _canaryRegion {
		return _stageReleaseTrain
	} else {
		return _stableReleaseTrain
	}
}
