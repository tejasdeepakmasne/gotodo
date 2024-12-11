package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tejasdeepakmasne/gotodo/operations"
)

var addCmd = &cobra.Command{
	Use:   "add \"description of task\" ",
	Short: "adds task to tasks.json",
	Long:  "adds task to the end of tasks.json",
	Run: func(cmd *cobra.Command, args []string) {
		operations.AddTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
