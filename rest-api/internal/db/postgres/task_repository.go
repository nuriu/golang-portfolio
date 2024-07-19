package postgres

import (
	"rest-api/internal/domain/task"

	"github.com/google/uuid"
)

// TODO: implement postgres repository
type TaskPostgresRepository struct{}

func NewTaskRepository() *TaskPostgresRepository {
	return &TaskPostgresRepository{}
}

func (repository *TaskPostgresRepository) Create(task *task.Task) error {
	return nil
}

func (repository *TaskPostgresRepository) List(page int, pageSize int) (*[]task.Task, error) {
	return nil, nil
}

func (repository *TaskPostgresRepository) Get(id uuid.UUID) (*task.Task, error) {
	return task.NewTask("title", "desc")
}

func (repository *TaskPostgresRepository) Delete(id uuid.UUID) error {
	return nil
}
