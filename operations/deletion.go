package operations

import "fmt"

func DeleteTask(id int) {
	tasks := LoadFile()

	if id < 1 || id > len(tasks) {
		fmt.Println("invalid task ID, valid IDs are:")
		ListTasks()
	} else {
		tasks = append(tasks[:id-1], tasks[id:]...)
		for newIDs := range tasks {
			tasks[newIDs].ID = newIDs + 1
		}
	}

	WriteTasks(&tasks)

}
