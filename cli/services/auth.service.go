package services

import (
	"encoding/json"
	"fmt"
	"pipebase/cli/types"
	"runtime"

	"github.com/99designs/keyring"
)

func SaveCredentials(userCredentials types.UserCredentials) {
	service := "pipebase"

	platform := runtime.GOOS

	var err error

	if platform == "darwin" || platform == "linux" {
		credentials, err = getCredentialsFromKeychain(service)
		if credentials != (types.UserCredentials{}) {
			fmt.Println("Pipebase administrator already exists")
			return
		}
	} else {
		// Handle other platforms (e.g., store in a file)
	}

	if err != nil {
		err = saveCredentialsToKeychain(userCredentials, service)
		if err != nil {
			fmt.Println("Error saving credentials to keychain:", err)
			return
		}
	}
}

func saveCredentialsToKeychain(credentials types.UserCredentials, service string) error {
	ring, err := keyring.Open(keyring.Config{
		ServiceName: service,
	})
	if err != nil {
		return err
	}

	credentialsJSON, err := json.Marshal(credentials)
	if err != nil {
		return err
	}

	if err := ring.Set(keyring.Item{Key: "credentials", Data: credentialsJSON}); err != nil {
		return err
	}

	fmt.Println("User credentials saved to keychain")

	return nil
}

func getCredentialsFromKeychain(service string) (types.UserCredentials, error) {
	var credentials types.UserCredentials

	ring, err := keyring.Open(keyring.Config{
		ServiceName: service,
	})
	if err != nil {
		return credentials, err
	}

	credentialsJSON, err := ring.Get("credentials")
	if err != nil {
		return credentials, err
	}

	err = json.Unmarshal([]byte(credentialsJSON.Data), &credentials)
	if err != nil {
		return credentials, err
	}

	return credentials, nil
}
