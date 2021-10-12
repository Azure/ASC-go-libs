package main

import (
	"fmt"
	"github.com/Azure/Tivan-Libs/pkg/config"
	"log"
	"os"
)

const (
	// _configFileKey is the key for the environment variable of the path to the config file.
	_configFileKey = "CONFIG_FILE"
)

type ExampleConfiguration struct {
	FirstName string
	LastName  string
}

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
	exampleConfiguration := new(ExampleConfiguration)

	/*Here you should add your configurations.*/
	keyConfigMap := map[string]interface{}{
		"exampleConfiguration": exampleConfiguration,
	}

	for key, configObject := range keyConfigMap {
		// Unmarshal the relevant parts of appConfig's data to each of the configuration objects
		err = config.CreateSubConfiguration(AppConfig, key, configObject)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to load specifc configuration data. \nkey: <%s>\nobjectType: <%T>", key, configObject)
			log.Fatal(errMsg)
		}
	}

	log.Print("Example configuration is ready for use: ", exampleConfiguration)
}
