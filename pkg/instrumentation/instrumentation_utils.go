package instrumentation

const (
	canaryRegion       = "eastus2euap"
	prodMdmAccount     = "RomeDetection"
	stageMdmAccount    = "RomeDetectionStage"
	stageReleaseTrain  = "stage"
	stableReleaseTrain = "stable"
)

func InitializeFromEnv(componentName, mdmNamespace string) *InstrumentationInitializationResult {
	configuration := NewInstrumentationConfigurationFromEnv(componentName, mdmNamespace)

	instrumentationInitializer := NewInstrumentationInitializer(configuration)
	instrumentationInitializationResults, err := instrumentationInitializer.Initialize()

	if err != nil {
		instrumentationInitializationResults.Tracer.Panic("error encountered during tracer initialization", err)
	}

	return instrumentationInitializationResults
}

func GetPlatformMdmAccount(region string) string {
	if region == canaryRegion {
		return stageMdmAccount
	} else {
		return prodMdmAccount
	}
}

func GetReleaseTrain(region string) string {
	if region == canaryRegion {
		return stageReleaseTrain
	} else {
		return stableReleaseTrain
	}
}
