package usecases_test

import (
	"errors"
	"testing"

	"github.com/NitzanShwartz/Task-Service/src/entities"
	usecases "github.com/NitzanShwartz/Task-Service/src/use_cases"
)

type TaskNotificationProviderMock struct {
	SendNotificationMock func(task entities.Task) error
}

func (tn *TaskNotificationProviderMock) SendNotification(task entities.Task) error {
	return tn.SendNotificationMock(task)
}

func TestNotifyTaskSuccess(t *testing.T) {
	notifierMock := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return nil
		},
	}

	useCase := usecases.NewNotifyTask(notifierMock)
	err := useCase.Execute(entities.Task{
		Title:       "test task",
		Description: "test task description",
		UserEmail:   "name@email.com",
	})
	if err != nil {
		t.Errorf("this function call should not fail, error: %v", err)
	}
}

func TestNotifyTaskFailed(t *testing.T) {
	notifierMock := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return errors.New("test error")
		},
	}

	useCase := usecases.NewNotifyTask(notifierMock)
	err := useCase.Execute(entities.Task{
		Title:       "test task",
		Description: "test task description",
		UserEmail:   "name@email.com",
	})
	if err.Error() != "test error" {
		t.Errorf("this function call should fail, and an error should be raised with the message 'test error', error: %v", err.Error())
	}
}
