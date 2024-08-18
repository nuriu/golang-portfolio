package task_test

import (
	"task-manager/internal/domain/task"
	"testing"
)

func TestNewTask(t *testing.T) {
	title := "title test"
	desc := "description test"

	t.Run("NewTask should not accept empty title", func(t *testing.T) {
		createdTask, err := task.NewTask("", desc)
		if err == nil {
			t.Error("NewTask should return error when received an empty title")
		}

		if createdTask != nil {
			t.Error("NewTask should not return task data when received an empty title")
		}

		if err != task.ErrorTaskTitleEmpty {
			t.Errorf("NewTask should return %s when received an empty title", task.ErrorTaskTitleEmpty.Error())
		}
	})

	t.Run("NewTask should not accept empty empty description", func(t *testing.T) {
		createdTask, err := task.NewTask(title, "")
		if err == nil {
			t.Error("NewTask should return error when received an empty description")
		}

		if createdTask != nil {
			t.Error("NewTask should not return task data when received an empty description")
		}

		if err != task.ErrorTaskDescriptionEmpty {
			t.Errorf("NewTask should return %s when received an empty description", task.ErrorTaskDescriptionEmpty.Error())
		}
	})

	t.Run("NewTask should create correct object", func(t *testing.T) {
		createdTask, err := task.NewTask(title, desc)
		if err != nil {
			t.Error("NewTask should not return error when received valid title and description")
		}

		if createdTask == nil {
			t.Error("NewTask should return task data when received valid title and description")
		} else if createdTask.Title != title {
			t.Errorf("NewTask should create task with correct title. Expected: %s, got: %s", title, createdTask.Title)
		} else if createdTask.Description != desc {
			t.Errorf("NewTask should create task with correct description. Expected: %s, got: %s", desc, createdTask.Description)
		}
	})
}
