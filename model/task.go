package task

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"desc"`
	CreatedAt   time.Time `json:"createdAt"`
	Status      bool      `json:"status"`
}

type AllTasks struct {
	Tasks []Task
}

func NewTask() Task {
	return Task{}
}
