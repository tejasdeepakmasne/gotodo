package operations

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mergestat/timediff"
)

func ListTasks() {
	tasks := LoadFile()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintf(writer, "ID\tTask\tCreated\tCompleted\t\n")
	for _, task := range tasks {
		friendlyTime := timediff.TimeDiff(task.CreatedAt)
		fmt.Fprintf(writer, "%d\t%s\t%s\t%t\t\n", task.ID, task.Description, friendlyTime, task.Status)
	}

	writer.Flush()
}
