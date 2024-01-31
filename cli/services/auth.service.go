package services

import (
	"cli/types"
	"encoding/json"
	"fmt"
	"os"
)

const credentialsEnvVar = "PIPEBASE_CREDENTIALS"

func SaveCredentials(userCredentials types.UserCredentials) error {
	// Check if credentials are already stored in the environment
	if os.Getenv(credentialsEnvVar) != "" {
		return fmt.Errorf("pipebase user already authenticated")
	}

	// Stringify the credentials
	credentialsJSON, err := json.Marshal(userCredentials)
	if err != nil {
		return err
	}

	// Store the credentials in the environment
	err = os.Setenv(credentialsEnvVar, string(credentialsJSON))
	if err != nil {
		return err
	}

	return nil
}

func GetCredentialsFromEnv() (types.UserCredentials, error) {
	// Retrieve credentials from the environment
	credentialsJSON := os.Getenv(credentialsEnvVar)
	if credentialsJSON == "" {
		return types.UserCredentials{}, fmt.Errorf("pipebase user not authenticated")
	}

	// Parse the JSON string to UserCredentials
	var credentials types.UserCredentials
	err := json.Unmarshal([]byte(credentialsJSON), &credentials)
	if err != nil {
		return types.UserCredentials{}, err
	}

	return credentials, nil
}
