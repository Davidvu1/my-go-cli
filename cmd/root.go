package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mygocli",
	Short: "davids own go cli",
}

func Execute() error {
	return rootCmd.Execute()
}
