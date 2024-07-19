package inmemory

import (
	"rest-api/internal/domain"
	"rest-api/internal/domain/task"

	"github.com/google/uuid"
)

var ErrorTaskNotFound = domain.NewDomainError(404, "task not found with given id")

type TaskInMemoryRepository struct {
	tasks []task.Task
}

func NewTaskRepository() *TaskInMemoryRepository {
	return &TaskInMemoryRepository{tasks: make([]task.Task, 10)}
}

func (repository *TaskInMemoryRepository) Create(task *task.Task) error {
	repository.tasks = appendTask(repository.tasks, *task)
	return nil
}

func (repository *TaskInMemoryRepository) List(page int, pageSize int) (*[]task.Task, error) {
	return &repository.tasks, nil
}

func (repository *TaskInMemoryRepository) Get(id uuid.UUID) (*task.Task, error) {
	for _, task := range repository.tasks {
		if task.ID == id {
			return &task, nil
		}

	}

	return nil, ErrorTaskNotFound
}

func (repository *TaskInMemoryRepository) Delete(id uuid.UUID) error {
	for i, task := range repository.tasks {
		if task.ID == id {
			repository.tasks[i] = repository.tasks[len(repository.tasks)-1]
			repository.tasks = repository.tasks[:len(repository.tasks)-1]
			return nil
		}
	}

	return ErrorTaskNotFound
}

func appendTask(tasks []task.Task, data ...task.Task) []task.Task {
	m := len(tasks)
	n := m + len(data)

	// if necessary, reallocate
	if n > cap(tasks) {
		// allocate double what's needed, for future growth.
		newSlice := make([]task.Task, (n+1)*2)
		copy(newSlice, tasks)
		tasks = newSlice
	}

	tasks = tasks[0:n]
	copy(tasks[m:n], data)
	return tasks
}
