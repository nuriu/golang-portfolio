package services_test

import (
	"errors"
	"task-manager/internal/app/services"
	"task-manager/internal/db/models"
	"task-manager/internal/db/repositories"
	"task-manager/internal/domain/task"
	"testing"
	"time"

	"github.com/google/uuid"
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

func TestGetTask(t *testing.T) {
	taskService, err := setupTakServiceTests()
	if err != nil {
		t.Error(err.Error())
	}

	t.Run("GetTask should return correct error when task not exists", func(t *testing.T) {
		id, _ := uuid.NewV7()
		taskDetail, err := taskService.GetTask(id)
		if err == nil {
			t.Error("GetUser should return correct error when task not exists")
		}

		if taskDetail != nil {
			t.Error("GetUser should not return any user data when task not exists")
		}

		if err != task.ErrorTaskNotFound {
			t.Errorf("GetUser should return correct error when task not exists")
		}
	})

	title, description := "title test", "description test"
	createdTask, _ := taskService.CreateTask(title, description)

	t.Run("GetTask should return correct task data", func(t *testing.T) {
		taskDetail, err := taskService.GetTask(createdTask.ID)
		if err != nil {
			t.Error("GetTask should not return any error when task exists")
		}

		if taskDetail != nil {
			if taskDetail.Title != createdTask.Title || taskDetail.Description != createdTask.Description {
				t.Error("GetTask should return correct task data")
			}
		}
	})
}

func TestListTasks(t *testing.T) {
	taskService, err := setupTakServiceTests()
	if err != nil {
		t.Error(err.Error())
	}

	title, description := "title test", "description test"
	title2, description2 := "title test 2", "description test 2"

	createdTask, _ := taskService.CreateTask(title, description)
	time.Sleep(500)
	createdTask2, _ := taskService.CreateTask(title2, description2)

	t.Run("ListTasks should return task list newer to older", func(t *testing.T) {
		tasks, err := taskService.ListTasks(1, 10)
		if err != nil {
			t.Error("ListTasks should not return error when there are tasks")
		}

		paginatedModel := tasks.(models.PaginatedModel)
		paginatedTasks := paginatedModel.Data.([]models.TaskEntity)

		for i, taskDetail := range paginatedTasks {
			if i == 0 && taskDetail.Title != createdTask2.Title {
				t.Error("error 1")
			}
			if i == 1 && taskDetail.Title != createdTask.Title {
				t.Error("error 2")
			}
		}
	})
}

func TestDeleteTask(t *testing.T) {
	taskService, err := setupTakServiceTests()
	if err != nil {
		t.Error(err.Error())
	}

	title, description := "title test", "description test"
	createdTask, _ := taskService.CreateTask(title, description)

	t.Run("DeleTask should delete the task with given ID", func(t *testing.T) {
		err := taskService.DeleteTask(createdTask.ID)
		if err != nil {
			t.Error("DeleTask should not return error when there are task to delete")
		}
	})
}
