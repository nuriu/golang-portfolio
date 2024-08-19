package task_test

import (
	"task-manager/internal/domain/task"
	"testing"
	"time"
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

	t.Run("NewTask should not accept empty description", func(t *testing.T) {
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

func TestUpdateTitle(t *testing.T) {
	title := "title test"
	desc := "description test"

	createdTask, err := task.NewTask(title, desc)
	if err != nil {
		t.Error("NewTask should not return error when received valid title and description")
	}

	lastUpdatedAt := createdTask.UpdatedAt

	err = createdTask.UpdateTitle("")
	if err == nil || err != task.ErrorTaskTitleEmpty {
		t.Error("UpdateTitle should return ErrorTaskTitleEmpty when received empty title")
	}

	time.Sleep(time.Millisecond * 1000)

	newTitle := "updated title test"
	err = createdTask.UpdateTitle(newTitle)
	if err != nil {
		t.Error("UpdateTitle should not return error when received valid title")
	}

	if createdTask.Title != newTitle {
		t.Errorf("UpdateTitle should update title correctly. Expected %s, got %s", newTitle, createdTask.Title)
	}

	if createdTask.UpdatedAt.Unix() <= lastUpdatedAt.Unix() {
		t.Errorf("UpdateTitle should also update UpdatedAt")
	}
}

func TestUpdateDescription(t *testing.T) {
	title := "title test"
	desc := "description test"

	createdTask, err := task.NewTask(title, desc)
	if err != nil {
		t.Error("NewTask should not return error when received valid title and description")
	}

	lastUpdatedAt := createdTask.UpdatedAt

	err = createdTask.UpdateDescription("")
	if err == nil || err != task.ErrorTaskDescriptionEmpty {
		t.Error("UpdateTitle should return ErrorTaskDescriptionEmpty when received empty description")
	}

	time.Sleep(time.Millisecond * 1000)

	newDescription := "updated description test"
	err = createdTask.UpdateDescription(newDescription)
	if err != nil {
		t.Error("UpdateTitle should not return error when received valid description")
	}

	if createdTask.Description != newDescription {
		t.Errorf("UpdateTitle should update description correctly. Expected %s, got %s", newDescription, createdTask.Description)
	}

	if createdTask.UpdatedAt.Unix() <= lastUpdatedAt.Unix() {
		t.Errorf("UpdateDescription should also update UpdatedAt")
	}
}

func TestCompletionOperations(t *testing.T) {
	title := "title test"
	desc := "description test"

	createdTask, err := task.NewTask(title, desc)
	if err != nil {
		t.Error("NewTask should not return error when received valid title and description")
	}

	if createdTask.CompletionDate != nil || createdTask.IsCompleted == true {
		t.Error("New tasks should created as not completed")
	}

	createdTask.Complete(time.Now().UTC())
	if createdTask.IsCompleted == false {
		t.Error("Complete should set task as completed")
	}

	createdTask.RevokeCompletion()
	if createdTask.IsCompleted == true || createdTask.CompletionDate != nil {
		t.Error("RevokeCompletion should set task as not completed")
	}
}
