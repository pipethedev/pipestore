/*
Copyright © 2024 NAME HERE davmuri1414@gmail.com
*/
package cmd

import (
	"cli/services"
	"cli/types"

	"github.com/spf13/cobra"
)

var username string
var passkey string

var config types.InitConfig

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup & Create pipebase administrative user for connection",
	Long:  `Setup your pipebase user, with a username and your connection api key is provided to you`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if username != "" {
			config.Username = username
		}
		if passkey != "" {
			config.PassKey = passkey
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		services.ExecuteInitialization(config, cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&username, "username", "u", "", "pipebase username")
	initCmd.Flags().StringVarP(&passkey, "passkey", "p", "", "pipebase passkey")
}
