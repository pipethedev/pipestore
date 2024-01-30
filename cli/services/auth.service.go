package services

import (
	"cli/types"
	"encoding/json"
	"fmt"
	"os"
)

func getCredentialsFilePath() string {
	return "/app/.pipebase_credentials"
}

func SaveCredentials(userCredentials types.UserCredentials) error {
	configFilePath := getCredentialsFilePath()

	if _, err := os.Stat(configFilePath); err == nil {
		return fmt.Errorf("pipebase user already authenticated")
	}

	err := saveCredentialsToConfigFile(userCredentials, configFilePath)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error saving credentials to config file")
	}

	return nil
}

func saveCredentialsToConfigFile(credentials types.UserCredentials, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	credentialsJSON, err := json.Marshal(credentials)
	if err != nil {
		return err
	}

	_, err = file.Write(credentialsJSON)
	if err != nil {
		return err
	}

	return nil
}
