package store

import (
	"errors"
	"github.com/google/uuid"
)

type Task struct {
	ID     uuid.UUID
	Title  string
	Status bool
}

type Store struct {
	db map[string]*Task
}

func NewStore() *Store {
	return &Store{
		db: make(map[string]*Task),
	}
}

func (s *Store) ListTasks() ([]*Task, error) {
	list := make([]*Task, 0, len(s.db))
	for _, task := range s.db {
		list = append(list, task)
	}

	return list, nil
}

func (s *Store) GetTaskByID(id uuid.UUID) (*Task, error) {
	if task, ok := s.db[id.String()]; ok {
		return task, nil
	}

	return nil, errors.New("task not found")
}

func (s *Store) CreateTask(title string) (*Task, error) {
	id := uuid.New()
	task := &Task{ID: id, Title: title}
	s.db[id.String()] = task

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
	task, err := s.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	delete(s.db, id.String())
	return task, nil
}
