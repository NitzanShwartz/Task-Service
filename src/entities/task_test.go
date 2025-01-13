package entities_test

import (
	"testing"

	"github.com/NitzanShwartz/Task-Service/src/entities"
)

func TestTaskCreationSuccess(t *testing.T) {
	task, err := entities.NewTask("test task", "test task description", "name@email.com")
	if err != nil {
		t.Errorf("this test should not fail")
	}

	if task.Title != "test task" || task.Description != "test task description" || task.UserEmail != "name@email.com" {
		t.Errorf("values were invalid for task instance %v", task)
	}
}

func TestTaskCreationFailNoTitle(t *testing.T) {
	task, err := entities.NewTask("", "test task description", "name@email.com")
	if err == nil {
		t.Errorf("this test should provide an error for missing value")
	}

	if task != nil {
		t.Errorf("task should be nil when this function fails")
	}
}

func TestTaskCreationFailNoDescription(t *testing.T) {
	task, err := entities.NewTask("task title", "", "name@email.com")
	if err == nil {
		t.Errorf("this test should provide an error for missing value")
	}

	if task != nil {
		t.Errorf("task should be nil when this function fails")
	}
}

func TestTaskCreationFailNoUserEmail(t *testing.T) {
	task, err := entities.NewTask("task title", "test task description", "")
	if err == nil {
		t.Errorf("this test should provide an error for missing value")
	}

	if task != nil {
		t.Errorf("task should be nil when this function fails")
	}
}
