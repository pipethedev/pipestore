package services

import (
	"fmt"
	"pipebase/cli/types"
	"runtime"

	"github.com/99designs/keyring"
)

func SaveCredentials(credentials types.UserCredentials) {
	service := "pipebase"

	platform := runtime.GOOS

	var err error

	if platform == "darwin" || platform == "linux" {
		credentials, err = getCredentialsFromKeychain(service)
	} else {
		// Handle other platforms (e.g., store in a file)
	}

	if err != nil {
		err = saveCredentialsToKeychain(credentials, service)
		if err != nil {
			fmt.Println("Error saving credentials to keychain:", err)
			return
		}
	}
}

func saveCredentialsToKeychain(credentials types.UserCredentials, service string) error {
	ring, _ := keyring.Open(keyring.Config{
		ServiceName: "pipebase_admin",
	})

	if err := ring.Set(keyring.Item{Key: "username", Data: []byte(string(credentials.Username))}); err != nil {
		return err
	}

	if err := ring.Set(keyring.Item{Key: "password", Data: []byte(string(credentials.Password))}); err != nil {
		return err
	}

	if err := ring.Set(keyring.Item{Key: "apiKey", Data: []byte(string(credentials.APIKey))}); err != nil {
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

	getItem := func(key string) (string, error) {
		item, err := ring.Get(key)
		if err != nil {
			return "", err
		}
		return string(item.Data), nil
	}

	credentials.Username, err = getItem("username")
	if err != nil {
		return credentials, err
	}

	credentials.Password, err = getItem("password")
	if err != nil {
		return credentials, err
	}

	credentials.APIKey, err = getItem("apiKey")
	if err != nil {
		return credentials, err
	}

	return credentials, nil
}
