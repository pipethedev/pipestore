package services

import (
	"cli/types"
	"cli/utils"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
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

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = "Installing PipeStore and starting container... "
	s.Start()

	err := InstallImageAndRunContainer()

	s.Stop()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Operation completed successfully ✅.")
}
