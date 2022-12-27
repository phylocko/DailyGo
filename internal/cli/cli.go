/*
Copyright © 2022 Vladislav Pavkin [phylocko@gmail.com]
*/
package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dailygo",
	Short: "Dailygo — a golang learning project.",
}

func init() {
	rootCmd.AddCommand(runserverCmd)

	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(initdbCmd)
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
