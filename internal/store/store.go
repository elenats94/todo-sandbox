package store

import (
	"errors"
	"github.com/google/uuid"
	"todo-sandbox/internal/models"
)

var errNotFound = errors.New("task not found")

type Store struct {
	db []*models.Task
}

func NewStore() *Store {
	return &Store{
		db: make([]*models.Task, 0),
	}
}

func (s *Store) ListTasks() ([]*models.Task, error) {
	return s.db, nil
}

func (s *Store) GetTaskByID(id uuid.UUID) (*models.Task, error) {
	for _, task := range s.db {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, errNotFound
}

func (s *Store) CreateTask(title string) (*models.Task, error) {
	id := uuid.New()
	task := &models.Task{ID: id, Title: title}
	s.db = append(s.db, task)

	return task, nil
}

func (s *Store) UpdateTask(id uuid.UUID, title string) (*models.Task, error) {
	task, err := s.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.Title = title
	return task, nil
}

func (s *Store) ToggleStatus(id uuid.UUID) (*models.Task, error) {
	task, err := s.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.Status = !task.Status
	return task, nil
}

func (s *Store) DeleteTask(id uuid.UUID) (*models.Task, error) {
	for i, task := range s.db {
		if task.ID == id {
			s.db = append(s.db[:i], s.db[i+1:]...)
			return task, nil
		}
	}

	return nil, errNotFound
}
