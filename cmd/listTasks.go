package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tejasdeepakmasne/gotodo/operations"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all the to-dos",
	Long:  "list all the available to-dos from tasks.json file in the current directort",
	Run: func(cmd *cobra.Command, args []string) {
		operations.ListTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
