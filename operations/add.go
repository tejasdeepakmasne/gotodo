package operations

import (
	"time"

	task "github.com/tejasdeepakmasne/gotodo/model"
)

func AddTask(description string) {
	tasks := LoadFile()
	newTask := task.Task{
		ID:          len(tasks) + 1,
		Description: description,
		CreatedAt:   time.Now(),
		Status:      false,
	}
	tasks = append(tasks, newTask)

	WriteTasks(&tasks)
}
