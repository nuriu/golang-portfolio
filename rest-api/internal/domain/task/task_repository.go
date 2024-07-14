package task

import (
	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(task *Task) (*Task, error)
	Get(id uuid.UUID) (*Task, error)
	Delete(id uuid.UUID) error
}
