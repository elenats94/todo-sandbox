package models

import "github.com/google/uuid"

type Task struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Status bool      `json:"status"`
}

func NewTask(title string) *Task {
	return &Task{
		ID:    uuid.New(),
		Title: title,
	}
}

type Tasks []*Task
