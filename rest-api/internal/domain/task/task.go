package task

import (
	"rest-api/internal/domain"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorTaskIDGeneration     = domain.NewDomainError(400, "task ID generation failed")
	ErrorTaskTitleEmpty       = domain.NewDomainError(400, "title is empty")
	ErrorTaskDescriptionEmpty = domain.NewDomainError(400, "description is empty")
)

type Task struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description string
	ID          uuid.UUID
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

	id, err := uuid.NewV7()
	if err != nil {
		return nil, ErrorTaskIDGeneration
	}

	return &Task{
		ID:          id,
		CreatedAt:   time.Now().UTC(),
		Title:       title,
		Description: description,
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
