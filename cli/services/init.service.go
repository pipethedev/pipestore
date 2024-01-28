package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"pipebase/cli/types"
	"pipebase/cli/utils"

	"github.com/spf13/cobra"
)

var credentials types.UserCredentials

// Todo: This should exist on a docker volume instead
const adminDirectory = ".pipebase/admin"

func ExecuteInitialization(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(adminDirectory); os.IsNotExist(err) {

		err := os.MkdirAll(adminDirectory, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating admin directory:", err)
			return
		}
	}

	credentialsFile := filepath.Join(adminDirectory, "pipestore_credentials.json")

	if _, err := os.Stat(credentialsFile); err == nil {
		// Credentials file already exists, throw an error
		fmt.Println("Pipebase administrator has been created before")
		return
	}

	if credentials.Username == "" {
		fmt.Print("Enter Pipebase username: ")
		fmt.Scanln(&credentials.Username)
	}

	if credentials.Password == "" {
		if key, err := utils.GenerateAPIKey(); err == nil {
			credentials.APIKey = key
		}
	}

	fmt.Printf("Updated credentials:\nUsername: %s\nAPI Key: %s\n", credentials.Username, credentials.APIKey)

	err := saveCredentialsToFile(credentials, credentialsFile)

	if err != nil {
		fmt.Println("Error saving credentials to file:", err)
		return
	}

	fmt.Println("Pipebase administrator created successfully.")
}

func saveCredentialsToFile(credentials types.UserCredentials, filePath string) error {
	encryptedCredentials, err := utils.Encrypt(credentials, "")

	if err != nil {
		return err
	}

	data, err := json.Marshal(encryptedCredentials)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, data, 0644)

	if err != nil {
		return err
	}

	return nil
}
