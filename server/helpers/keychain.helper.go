package helpers

import (
	"server/types"
)

// const credentialsEnvVar = "PIPEBASE_CREDENTIALS"

func GetCredentialsFromEnv() (types.UserCredentials, error) {
	return types.UserCredentials{
		Username:    "pipethedev",
		Password:    "passkey",
		APIKey:      "7yCeF2uY7pfhgMEOZGy42CvEca2AfD18bwEOS2FBWdA=",
		StoragePath: "",
	}, nil
}
