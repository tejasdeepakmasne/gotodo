package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tejasdeepakmasne/gotodo/operations"
)

var markDone = &cobra.Command{
	Use:   "done <task ID>",
	Short: "marks the task as done",
	Long:  "marks the status of task completion with the ID in tasks.json as true",
	Run: func(cmd *cobra.Command, args []string) {
		argInInt, _ := strconv.Atoi(args[0])
		operations.MarkDone(argInInt)
	},
}

func init() {
	rootCmd.AddCommand(markDone)
}
