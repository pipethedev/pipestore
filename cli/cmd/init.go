/*
Copyright © 2024 NAME HERE davmuri1414@gmail.com
*/
package cmd

import (
	"pipebase/cli/services"
	"pipebase/cli/types"

	"github.com/spf13/cobra"
)

var config types.InitConfig

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create pipebase administrative user for connection",
	Long:  `Setup your pipebase user, with a username and your connection api key is provided to you`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if username != "" {
			config.Username = username
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		services.ExecuteInitialization(config, cmd, args)
	},
}

var username string

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&username, "username", "u", "", "pipebase username")
	initCmd.Flags().StringVarP(&username, "passkey", "p", "", "pipebase passkey")
}
