package task

import (
	"strings"
	"task-manager/internal/domain"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorTaskIDGeneration     = domain.NewDomainError(400, "task ID generation failed")
	ErrorTaskTitleEmpty       = domain.NewDomainError(400, "title is empty")
	ErrorTaskDescriptionEmpty = domain.NewDomainError(400, "description is empty")
)

type Task struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Title          string
	Description    string
	ID             uuid.UUID
	IsCompleted    bool
	CompletionDate *time.Time
}

func NewTask(title string, description string) (*Task, error) {
	title = strings.Trim(title, " ")
	if len(title) < 1 {
		return nil, ErrorTaskTitleEmpty
	}

	description = strings.Trim(description, " ")
	if len(description) < 1 {
		return nil, ErrorTaskDescriptionEmpty
	}

	return &Task{
		ID:             uuid.Nil,
		Title:          title,
		Description:    description,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		DeletedAt:      nil,
		IsCompleted:    false,
		CompletionDate: nil,
	}, nil
}

func (task *Task) UpdateTitle(title string) error {
	title = strings.Trim(title, " ")
	if len(title) < 1 {
		return ErrorTaskTitleEmpty
	}

	task.Title = title
	task.UpdatedAt = time.Now().UTC()

	return nil
}

func (task *Task) UpdateDescription(description string) error {
	description = strings.Trim(description, " ")
	if len(description) < 1 {
		return ErrorTaskDescriptionEmpty
	}

	task.Description = description
	task.UpdatedAt = time.Now().UTC()

	return nil
}

func (task *Task) Complete(completionDate time.Time) {
	task.IsCompleted = true
	task.CompletionDate = &completionDate
}

func (task *Task) RevokeCompletion() {
	task.IsCompleted = false
	task.CompletionDate = nil
}
