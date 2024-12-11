package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tejasdeepakmasne/gotodo/operations"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <task-ID>",
	Short: "deletes the task with the given ID",
	Long:  "deletes the task with the given ID from tasks.json",
	Run: func(cmd *cobra.Command, args []string) {
		argInInt, _ := strconv.Atoi(args[0])
		operations.DeleteTask(argInInt)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
