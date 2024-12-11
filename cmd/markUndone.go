package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tejasdeepakmasne/gotodo/operations"
)

var markUnDone = &cobra.Command{
	Use:   "revert <task ID>",
	Short: "marks the task as not done",
	Long:  "marks the status of task completion with the ID in tasks.json as false",
	Run: func(cmd *cobra.Command, args []string) {
		argInInt, _ := strconv.Atoi(args[0])
		operations.MarkUndone(argInInt)
	},
}

func init() {
	rootCmd.AddCommand(markUnDone)
}
