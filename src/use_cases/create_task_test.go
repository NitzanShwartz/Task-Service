package usecases_test

import (
	"errors"
	"testing"

	"github.com/NitzanShwartz/Task-Service/src/entities"
	"github.com/NitzanShwartz/Task-Service/src/use_cases"
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

func TestCreateTaskSuccess(t *testing.T) {
	repositoryMock := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return false
		},
		CreateTaskMock: func(task entities.Task) error {
			return nil
		},
	}

	notifierMock := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return nil
		},
	}

	useCase := usecases.NewCreateTask(repositoryMock, notifierMock)
	err := useCase.Execute("test task", "test task description", "name@email.com")
	if err != nil {
		t.Error("this function call should not fail")
	}
}

func TestCreateTaskFailNoTaskName(t *testing.T) {
	repositoryMock := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return false
		},
		CreateTaskMock: func(task entities.Task) error {
			return nil
		},
	}

	notifierMock := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return nil
		},
	}

	useCase := usecases.NewCreateTask(repositoryMock, notifierMock)
	err := useCase.Execute("", "test task description", "name@email.com")
	if err == nil {
		t.Error("this function call should fail, missing name for task!")
	}
}

func TestCreateTaskFailNoTaskDescription(t *testing.T) {
	repositoryMock := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return false
		},
		CreateTaskMock: func(task entities.Task) error {
			return nil
		},
	}

	notifierMock := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return nil
		},
	}

	useCase := usecases.NewCreateTask(repositoryMock, notifierMock)
	err := useCase.Execute("test task", "", "name@email.com")
	if err == nil {
		t.Error("this function call should fail, missing name for task!")
	}
}

func TestCreateTaskFailNoTaskUserEmail(t *testing.T) {
	repositoryMock := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return false
		},
		CreateTaskMock: func(task entities.Task) error {
			return nil
		},
	}

	notifierMock := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return nil
		},
	}

	useCase := usecases.NewCreateTask(repositoryMock, notifierMock)
	err := useCase.Execute("test task", "test task description", "")
	if err == nil {
		t.Error("this function call should fail, missing name for task!")
	}
}

func TestCreateTaskFailTaskUserEmailDoesNotMatchConvention(t *testing.T) {
	repositoryMock := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return false
		},
		CreateTaskMock: func(task entities.Task) error {
			return nil
		},
	}

	notifierMock := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return nil
		},
	}

	useCase := usecases.NewCreateTask(repositoryMock, notifierMock)
	err := useCase.Execute("test task", "test task description", "nameemail.com")
	if err == nil {
		t.Error("this function call should fail, missing name for task!")
	}
}

func TestCreateTaskFailTaskRepositoryCreateFailed(t *testing.T) {
	repositoryMock := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return false
		},
		CreateTaskMock: func(task entities.Task) error {
			return errors.New("test error")
		},
	}

	notifierMock := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return nil
		},
	}

	useCase := usecases.NewCreateTask(repositoryMock, notifierMock)
	err := useCase.Execute("test task", "test task description", "name@email.com")
	if err.Error() != "test error" {
		t.Errorf("this function call should fail, and an error should be raised with the message 'test error', error: %v", err.Error())
	}
}

func TestCreateTaskFailTaskRepositoryTaskAlreadyExists(t *testing.T) {
	repositoryMock := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return true
		},
		CreateTaskMock: func(task entities.Task) error {
			return nil
		},
	}

	notifierMock := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return nil
		},
	}

	useCase := usecases.NewCreateTask(repositoryMock, notifierMock)
	err := useCase.Execute("test task", "test task description", "name@email.com")
	var tae *exceptions.TaskAlreadyExistsError
	if !errors.As(err, &tae) {
		t.Errorf("this function call should fail, and an error should be raised with the message 'test error', error: %v", err.Error())
	}
}

func TestCreateTaskFailTaskNotificationFailed(t *testing.T) {
	repositoryMock := &TaskRepositoryMock{
		DoesTaskExistsMock: func(title string) bool {
			return false
		},
		CreateTaskMock: func(task entities.Task) error {
			return nil
		},
	}

	notifierMock := &TaskNotificationProviderMock{
		SendNotificationMock: func(task entities.Task) error {
			return errors.New("test error")
		},
	}

	useCase := usecases.NewCreateTask(repositoryMock, notifierMock)
	err := useCase.Execute("test task", "test task description", "name@email.com")
	if err.Error() != "test error" {
		t.Errorf("this function call should fail, and an error should be raised with the message 'test error', error: %v", err.Error())
	}
}
