package store

import (
	"errors"
	"github.com/google/uuid"
	"todo-sandbox/internal/models"
)

var errNotFound = errors.New("task not found")

type Store struct {
	db map[uuid.UUID]models.Tasks
}

func NewStore() *Store {
	return &Store{
		db: make(map[uuid.UUID]models.Tasks),
	}
}

func (s *Store) ListTasks(owner uuid.UUID) (models.Tasks, error) {
	return s.db[owner], nil
}

func (s *Store) GetTaskByID(taskID, owner uuid.UUID) (*models.Task, error) {
	tasks, err := s.ListTasks(owner)
	if err != nil {
		return nil, errNotFound
	}

	for _, task := range tasks {
		if task.ID == taskID {
			return task, nil
		}
	}

	return nil, errNotFound
}

func (s *Store) CreateTask(title string, owner uuid.UUID) (*models.Task, error) {
	task := models.NewTask(title)
	s.db[owner] = append(s.db[owner], task)

	return task, nil
}

func (s *Store) UpdateTask(id uuid.UUID, title string, owner uuid.UUID) (*models.Task, error) {
	task, err := s.GetTaskByID(id, owner)
	if err != nil {
		return nil, err
	}

	task.Title = title
	return task, nil
}

func (s *Store) ToggleStatus(id uuid.UUID, owner uuid.UUID) (*models.Task, error) {
	task, err := s.GetTaskByID(id, owner)
	if err != nil {
		return nil, err
	}

	task.Status = !task.Status
	return task, nil
}

func (s *Store) DeleteTask(id uuid.UUID, owner uuid.UUID) (*models.Task, error) {
	tasks, err := s.ListTasks(owner)
	if err != nil {
		return nil, err
	}

	for i, task := range tasks {
		if task.ID == id {
			s.db[owner] = append(tasks[:i], tasks[i+1:]...)
			return task, nil
		}
	}

	return nil, errNotFound
}
