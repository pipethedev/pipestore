package helpers

import (
	"encoding/json"
	"server/types"

	"github.com/99designs/keyring"
)

func GetCredentialsFromKeychain(service string) (types.UserCredentials, error) {
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
