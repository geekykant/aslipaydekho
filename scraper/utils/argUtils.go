package utils

import "os"

func CheckAllEnvVarsPresent() bool {
	envVars := []string{"RABBIT_MQ_SERVER_URL", "RABBIT_MQ_CHANNEL_NAME", "RABBITMQ_USERNAME", "RABBITMQ_PASSWORD"}
	for _, varKey := range envVars {
		_, isSet := os.LookupEnv(varKey)
		if !isSet {
			return false
		}
	}
	return true
}
