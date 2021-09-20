package common

import (
	"os"
)

func GetEnvVariableOrDefault(key string, defaultValue string) string {
	value, keyExists := os.LookupEnv(key)

	if !keyExists {
		return defaultValue
	}

	return value
}
