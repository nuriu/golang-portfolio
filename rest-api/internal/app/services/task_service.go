package services

import (
	"rest-api/internal/domain/task"
)

type TaskService struct {
	taskRepository task.TaskRepository
}

func NewTaskService(taskRepository task.TaskRepository) task.TaskService {
	return &TaskService{taskRepository}
}

func (service *TaskService) CreateTask(title string, description string) (*task.Task, error) {
	task, err := task.NewTask(title, description)
	return task, err
}
