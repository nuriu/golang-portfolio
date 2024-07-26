package services

import (
	"task-manager/internal/domain/task"

	"github.com/google/uuid"
)

type TaskService struct {
	taskRepository task.TaskRepository
}

func NewTaskService(taskRepository task.TaskRepository) task.TaskService {
	return &TaskService{taskRepository}
}

// CreateTask implements task.TaskService.
func (t *TaskService) CreateTask(title string, description string) (*task.Task, error) {
	generatedTask, err := task.NewTask(title, description)
	if err != nil {
		return nil, err
	}

	createdTask, err := t.taskRepository.Create(generatedTask)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

// ListTasks implements task.TaskService.
func (t *TaskService) ListTasks(page int, pageSize int) (interface{}, error) {
	tasks, err := t.taskRepository.List(page, pageSize)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTask implements task.TaskService.
func (t *TaskService) GetTask(taskID uuid.UUID) (*task.Task, error) {
	taskDetail, err := t.taskRepository.Get(taskID)
	if err != nil {
		return nil, err
	}

	return taskDetail, nil
}

// DeleteTask implements task.TaskService.
func (t *TaskService) DeleteTask(taskID uuid.UUID) error {
	err := t.taskRepository.Delete(taskID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTask implements task.TaskService.
func (t *TaskService) UpdateTask(taskID uuid.UUID, title string, description string) error {
	// TODO: update to user taskRepository
	panic("unimplemented")
}
