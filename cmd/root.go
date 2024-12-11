package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gotodo",
	Short: "gotodo is a cli to-do list application",
	Long:  `gotodo is a simple cli to-do list application which leverages a tasks.json file in the current directory so that you can have multiple such lists across your different projects`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
