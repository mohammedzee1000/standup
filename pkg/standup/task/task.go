package task

import "github.com/google/uuid"

type Task struct {
	ID          string `json:"ID"`
	Description string `json:"Description"`
}

func NewEmptyTask() *Task {
	return &Task{}
}

func New() *Task {
	t := NewEmptyTask()
	t.ID = uuid.New().String()
	return t
}
