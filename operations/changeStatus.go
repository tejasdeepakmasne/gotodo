package operations

import (
	"fmt"
)

func MarkDone(id int) {
	tasks := LoadFile()

	if id < 1 || id > len(tasks) {
		fmt.Println("invalid id, these are the valid IDs")
		ListTasks()
	} else {
		tasks[id-1].Status = true
	}

	WriteTasks(&tasks)

}

func MarkUndone(id int) {
	tasks := LoadFile()

	if id < 1 || id > len(tasks) {
		fmt.Println("invalid id, these are the valid IDs")
		ListTasks()
	} else {
		tasks[id-1].Status = false
	}

	WriteTasks(&tasks)
}
