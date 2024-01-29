package services

import (
	"fmt"
	"pipebase/cli/types"
	"pipebase/cli/utils"

	"github.com/spf13/cobra"
)

var credentials types.UserCredentials

func ExecuteInitialization(config types.InitConfig, cmd *cobra.Command, args []string) {
	credentials.Username = config.Username
	credentials.Password = config.PassKey

	if credentials.Username == "" {
		fmt.Print("Enter Pipebase username: ")
		fmt.Scanln(&credentials.Username)
	}

	if credentials.Password == "" {
		fmt.Print("Enter Pipebase passkey: ")
		fmt.Scanln(&credentials.Password)
	}

	if key, err := utils.GenerateAPIKey(); err == nil {
		credentials.APIKey = key
	}

	SaveCredentials(credentials)

	fmt.Println("Pipebase administrator created successfully.")
}

// func saveCredentialsToFile(credentials types.UserCredentials, filePath string, passKey string) error {
// 	encryptedCredentials, err := utils.Encrypt(credentials, passKey)

// 	if err != nil {
// 		return err
// 	}

// 	data, err := json.Marshal(encryptedCredentials)

// 	if err != nil {
// 		return err
// 	}

// 	err = ioutil.WriteFile(filePath, data, 0644)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
