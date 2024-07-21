package task

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(title string, description string) error
	ListTasks(page int, pageSize int) (interface{}, error)
	GetTask(taskID uuid.UUID) (*Task, error)
	UpdateTask(taskID uuid.UUID, title string, description string) error
	DeleteTask(taskID uuid.UUID) error
}
