package models

import (
	"rest-api/internal/domain/task"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description string
	ID          uuid.UUID `gorm:"primaryKey"`
}

func MapToDBTask(task *task.Task) *Task {
	return &Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func MapFromDBTask(dbTask *Task) *task.Task {
	return &task.Task{
		ID:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description,
		CreatedAt:   dbTask.CreatedAt,
		UpdatedAt:   dbTask.UpdatedAt,
	}
}
