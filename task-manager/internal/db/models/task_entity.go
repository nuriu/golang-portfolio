package models

import (
	"task-manager/internal/domain/task"
)

type TaskEntity struct {
	BaseEntity
	Title       string
	Description string
}

func (TaskEntity) TableName() string {
	return "tasks"
}

func FromDomain(t *task.Task) *TaskEntity {
	return &TaskEntity{
		BaseEntity: BaseEntity{
			ID:        t.ID,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			DeletedAt: t.DeletedAt,
		},
		Title:       t.Title,
		Description: t.Description,
	}
}

func (te *TaskEntity) ToDomain() *task.Task {
	return &task.Task{
		ID:          te.ID,
		Title:       te.Title,
		Description: te.Description,
		CreatedAt:   te.CreatedAt,
		UpdatedAt:   te.UpdatedAt,
		DeletedAt:   te.DeletedAt,
	}
}
