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
	task, err := task.NewTask(title, description)
	if err != nil {
		return err
	}

	if err := t.taskRepository.Create(task); err != nil {
		return err
	}

	return nil
}

// ListTasks implements task.TaskService.
func (t *TaskService) ListTasks(page int, pageSize int) ([]*task.Task, error) {
	tasks, err := t.taskRepository.List(1, 10)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTask implements task.TaskService.
func (t *TaskService) GetTask(taskID uuid.UUID) (*task.Task, error) {
	task, err := t.taskRepository.Get(taskID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// DeleteTask implements task.TaskService.
func (t *TaskService) DeleteTask(taskID uuid.UUID) error {
	if err := t.taskRepository.Delete(taskID); err != nil {
		return err
	}

	return nil
}

// UpdateTask implements task.TaskService.
func (t *TaskService) UpdateTask(taskID uuid.UUID, title string, description string) error {
	// TODO: update to user taskRepository
	panic("unimplemented")
}
