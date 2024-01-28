/*
Copyright © 2024 NAME HERE davmuri1414@gmail.com
*/
package cmd

import (
	"pipebase/cli/services"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create pipebase administrative user for connection",
	Long:  `Setup your pipebase user, with a username and your connection api key is provided to you`,
	Run:   services.ExecuteInitialization,
}

var username string

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&username, "username", "u", "", "pipebase username")
}
