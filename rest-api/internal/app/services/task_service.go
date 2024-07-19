package services

import (
	"rest-api/internal/domain/task"

	"github.com/google/uuid"
)

type TaskService struct {
	taskRepository task.TaskRepository
}

func NewTaskService(taskRepository task.TaskRepository) task.TaskService {
	return &TaskService{taskRepository}
}

// CreateTask implements task.TaskService.
func (t *TaskService) CreateTask(title string, description string) error {
	// TODO: update to user taskRepository
	_, err := task.NewTask(title, description)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTask implements task.TaskService.
func (t *TaskService) DeleteTask(taskID uuid.UUID) error {
	// TODO: update to user taskRepository
	panic("unimplemented")
}

// ListTasks implements task.TaskService.
func (t *TaskService) ListTasks(page int, pageSize int) (*[]task.Task, error) {
	// TODO: update to user taskRepository
	panic("unimplemented")
}

// UpdateTask implements task.TaskService.
func (t *TaskService) UpdateTask(taskID uuid.UUID, title string, description string) error {
	// TODO: update to user taskRepository
	panic("unimplemented")
}
