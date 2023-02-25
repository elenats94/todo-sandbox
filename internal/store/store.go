package store

import (
	"errors"
	"github.com/google/uuid"
)

var errNotFound = errors.New("task not found")

type Task struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Status bool      `json:"status"`
}

type Store struct {
	db []*Task
}

func NewStore() *Store {
	return &Store{
		db: make([]*Task, 0),
	}
}

func (s *Store) ListTasks() ([]*Task, error) {
	return s.db, nil
}

func (s *Store) GetTaskByID(id uuid.UUID) (*Task, error) {
	for _, task := range s.db {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, errNotFound
}

func (s *Store) CreateTask(title string) (*Task, error) {
	id := uuid.New()
	task := &Task{ID: id, Title: title}
	s.db = append(s.db, task)

	return task, nil
}

func (s *Store) UpdateTask(id uuid.UUID, title string) (*Task, error) {
	task, err := s.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.Title = title
	return task, nil
}

func (s *Store) ToggleStatus(id uuid.UUID) (*Task, error) {
	task, err := s.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.Status = !task.Status
	return task, nil
}

func (s *Store) DeleteTask(id uuid.UUID) (*Task, error) {
	for i, task := range s.db {
		if task.ID == id {
			s.db = append(s.db[:i], s.db[i+1:]...)
			return task, nil
		}
	}

	return nil, errNotFound
}
