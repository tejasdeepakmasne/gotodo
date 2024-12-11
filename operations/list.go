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
	fmt.Fprintf(writer, "ID\tTask\tCreated\tStatus\t\n")
	for _, task := range tasks {
		friendlyTime := timediff.TimeDiff(task.CreatedAt)
		if task.Status {
			fmt.Fprintf(writer, "%d\t%s\t%s\tdone\t\n", task.ID, task.Description, friendlyTime)
		} else {
			fmt.Fprintf(writer, "%d\t%s\t%s\tnot done\t\n", task.ID, task.Description, friendlyTime)
		}

	}

	writer.Flush()
}
