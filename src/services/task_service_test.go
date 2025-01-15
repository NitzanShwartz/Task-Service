package services_test

import (
	"errors"
	"testing"

	"github.com/NitzanShwartz/Task-Service/src/entities"
	"github.com/NitzanShwartz/Task-Service/src/services"
	"github.com/NitzanShwartz/Task-Service/src/use_cases/exceptions"
)

type TaskRepositoryMock struct {
	DoesTaskExistsMock func(title string) bool
	CreateTaskMock     func(task entities.Task) error
}

func (trm *TaskRepositoryMock) DoesTaskExists(title string) bool {
	return trm.DoesTaskExistsMock(title)
}

func (tr *TaskRepositoryMock) CreateTask(task entities.Task) error {
	return tr.CreateTaskMock(task)
}

type TaskNotificationProviderMock struct {
	SendNotificationMock func(task entities.Task) error
}

func (tn *TaskNotificationProviderMock) SendNotification(task entities.Task) error {
	return tn.SendNotificationMock(task)
}

func TestTaskServiceCreateTaskSuccess(t *testing.T) {
	taskRepository := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return false
		},
		CreateTaskMock: func(task entities.Task) error {
			return nil
		},
	}

	notificationRepository := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return nil
		},
	}

	taskService := services.NewTaskService(taskRepository, notificationRepository)
	err := taskService.CreateTask("test task", "test task description", "name@email.com")
	if err != nil {
		t.Errorf("this function call should not fail, error: %v", err)
	}
}

func TestTaskServiceCreateTasKFailkTaskExists(t *testing.T) {
	taskRepository := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return true
		},
		CreateTaskMock: func(task entities.Task) error {
			return nil
		},
	}

	notificationRepository := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return nil
		},
	}

	taskService := services.NewTaskService(taskRepository, notificationRepository)
	err := taskService.CreateTask("test task", "test task description", "name@email.com")
	if err == nil {
		t.Error("this function call should fail")
	}

    var tae *exceptions.TaskAlreadyExistsError
    if !errors.As(err, &tae) {
        t.Error("not the correct error")
    }
}

func TestTaskServiceCreateTasKFailkCreateFailed(t *testing.T) {
	testErrMsg := "test error - createTaskMock"
	taskRepository := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return false
		},
		CreateTaskMock: func(task entities.Task) error {
			return errors.New(testErrMsg)
		},
	}

	notificationRepository := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return nil
		},
	}

	taskService := services.NewTaskService(taskRepository, notificationRepository)
	err := taskService.CreateTask("test task", "test task description", "name@email.com")
	if err == nil {
		t.Error("this function call should fail")
	}
	if err.Error() != testErrMsg {
		t.Errorf("the error that was raised is not the expected error, error: %v", err)
	}
}

func TestTaskServiceCreateTasKFailkNotifyFailed(t *testing.T) {
	testErrMsg := "test error - notificationRepistoryMock"
	taskRepository := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return false
		},
		CreateTaskMock: func(task entities.Task) error {
			return nil
		},
	}

	notificationRepository := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return errors.New(testErrMsg)
		},
	}

	taskService := services.NewTaskService(taskRepository, notificationRepository)
	err := taskService.CreateTask("test task", "test task description", "name@email.com")
	if err == nil {
		t.Error("this function call should fail")
	}
	if err.Error() != testErrMsg {
		t.Errorf("the error that was raised is not the expected error, error: %v", err)
	}
}
