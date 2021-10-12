package main

import (
	"fmt"
	"github.com/Azure/Tivan-Libs/pkg/config"
	"github.com/Azure/Tivan-Libs/pkg/instrumentation"
	"log"
	"os"
)

const (
	// _configFileKey is the key for the environment variable of the path to the config file.
	_configFileKey = "CONFIG_FILE"
)

func main() {

	configFilePath := os.Getenv(_configFileKey)
	if len(configFilePath) == 0 {
		log.Fatalf("%v env variable is not defined.", _configFileKey)
	}
	// Load configuration
	AppConfig, err := config.LoadConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Create Configuration objects for each factory
	instrumenationConfiguration := new(instrumentation.InstrumentationConfiguration)

	keyConfigMap := map[string]interface{}{
		"instrumentation": instrumenationConfiguration,
		/*Here you should add your configurations.*/
	}

	for key, configObject := range keyConfigMap {
		// Unmarshal the relevant parts of appConfig's data to each of the configuration objects
		err = config.CreateSubConfiguration(AppConfig, key, configObject)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to load specifc configuration data. \nkey: <%s>\nobjectType: <%T>", key, configObject)
			log.Fatal(errMsg)
		}
	}

	initializer := instrumentation.NewInstrumentationInitializer(instrumenationConfiguration)
	instumentation, err := initializer.Initialize()
	if err != nil {
		log.Fatal(err, "failed to initialize instrumentation")
	}
	/*The instrumentation is ready for use.*/
	_ = instumentation.Tracer
}
