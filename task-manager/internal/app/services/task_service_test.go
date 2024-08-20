package services_test

import (
	"errors"
	"task-manager/internal/app/services"
	"task-manager/internal/db/repositories"
	"task-manager/internal/domain/task"
	"testing"

	gormSqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTakServiceTests() (task.TaskService, error) {
	db, err := gorm.Open(gormSqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to open sqlite db")
	}

	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)

	return taskService, nil
}

func TestCreateTask(t *testing.T) {
	taskService, err := setupTakServiceTests()
	if err != nil {
		t.Error(err.Error())
	}

	title, description := "title test", "description test"
	t.Run("CreateTask should be able to create a new task", func(t *testing.T) {
		createdTask, err := taskService.CreateTask(title, description)
		if err != nil {
			t.Error("CreateTask should not return any errors when received valid data")
		}

		if createdTask == nil {
			t.Error("CreateTask should return created task")

		}
	})

	t.Run("CreateTask should not create a task with empty title", func(t *testing.T) {
		createdTask, err := taskService.CreateTask("", description)

		if err == nil {
			t.Error("CreateTask should return error when received an empty title")
		}

		if createdTask != nil {
			t.Error("CreateUser should not return any task data when received an empty title")
		}

		if err != task.ErrorTaskTitleEmpty {
			t.Errorf("CreateUser should return %s when received an empty title", task.ErrorTaskTitleEmpty.Error())
		}
	})

	t.Run("CreateTask should not create a task with empty description", func(t *testing.T) {
		createdTask, err := taskService.CreateTask(title, "")

		if err == nil {
			t.Error("CreateTask should return error when received an empty description")
		}

		if createdTask != nil {
			t.Error("CreateUser should not return any task data when received an empty description")
		}

		if err != task.ErrorTaskDescriptionEmpty {
			t.Errorf("CreateUser should return %s when received an empty description", task.ErrorTaskDescriptionEmpty.Error())
		}
	})
}
