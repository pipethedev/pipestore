package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"server/types"
)

func GetCredentialsFromStore() (types.UserCredentials, error) {
	configFilePath := "/app/.pipebase_credentials"

	file, err := os.Open(configFilePath)
	if err != nil {
		return types.UserCredentials{}, fmt.Errorf("error opening credentials file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var credentials types.UserCredentials
	err = decoder.Decode(&credentials)
	if err != nil {
		return types.UserCredentials{}, fmt.Errorf("error decoding credentials")
	}

	return credentials, nil
}
