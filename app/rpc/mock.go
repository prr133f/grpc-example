package rpc

import (
	"todo/app/models"

	"github.com/google/uuid"
)

type mock struct {
}

func (m *mock) CreateTask(title string, description string) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (m *mock) ResolveTask(id uuid.UUID) error {
	return nil
}

func (m *mock) GetTaskList() ([]models.Task, error) {
	return nil, nil
}

func (m *mock) GetTaskById(id uuid.UUID) (models.Task, error) {
	return models.Task{}, nil
}

func (m *mock) DeleteTask(id uuid.UUID) error {
	return nil
}
