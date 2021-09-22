package common

import (
	"os"
)

// GetEnvVariableOrDefault this method gets a key and default value.
// If the key is env variable, it returns the value of the env.
// If the key is not env variable, it returns the default value.
func GetEnvVariableOrDefault(key string, defaultValue string) string {
	value, keyExists := os.LookupEnv(key)

	if !keyExists {
		return defaultValue
	}

	return value
}
