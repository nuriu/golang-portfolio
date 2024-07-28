package models

import (
	"task-manager/internal/domain/task"

	"gorm.io/gorm"
)

type TaskEntity struct {
	BaseEntity
	Title       string
	Description string
}

func (TaskEntity) TableName() string {
	return "tasks"
}

func TaskFromDomainEntity(t *task.Task) *TaskEntity {
	deleteInfo := gorm.DeletedAt{Valid: false}
	if t.DeletedAt != nil {
		deleteInfo.Time = *t.DeletedAt
		deleteInfo.Valid = true
	}

	return &TaskEntity{
		BaseEntity: BaseEntity{
			ID:        t.ID,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			DeletedAt: deleteInfo,
		},
		Title:       t.Title,
		Description: t.Description,
	}
}

func (te *TaskEntity) ToDomainEntity() *task.Task {
	return &task.Task{
		ID:          te.ID,
		Title:       te.Title,
		Description: te.Description,
		CreatedAt:   te.CreatedAt,
		UpdatedAt:   te.UpdatedAt,
		DeletedAt:   &te.DeletedAt.Time,
	}
}
