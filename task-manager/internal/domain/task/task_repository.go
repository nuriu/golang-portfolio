package task

import (
	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(task *Task) (*Task, error)
	List(page int, pageSize int) (interface{}, error)
	Get(id uuid.UUID) (*Task, error)
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, task *Task)
}
