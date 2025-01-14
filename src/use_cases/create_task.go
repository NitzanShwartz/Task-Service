package usecases

import (
	"errors"
	"fmt"

	"regexp"

	"github.com/NitzanShwartz/Task-Service/src/entities"
	"github.com/NitzanShwartz/Task-Service/src/repositories"
	"github.com/NitzanShwartz/Task-Service/src/use_cases/exceptions"
)

const emailRegExp = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"

type CreateTask struct {
	TaskRepository       repositories.TaskRepository
	NotificationProvider repositories.TaskNotificationProvider
}

func NewCreateTask(taskRepsitory repositories.TaskRepository, notificationProvider repositories.TaskNotificationProvider) *CreateTask {
	return &CreateTask{
		TaskRepository:       taskRepsitory,
		NotificationProvider: notificationProvider,
	}
}

func (ct *CreateTask) Execute(title string, description string, userEmail string) error {
	if title == "" || description == "" || userEmail == "" {
		return errors.New("title, task and userEmail are mandatory fields")
	}

	if match, _ := regexp.MatchString(emailRegExp, userEmail); !match {
		return errors.New("emails must be of the format <name>@<domain.tld>")
	}

	if ct.TaskRepository.DoesTaskExists(title) {
		return &exceptions.TaskAlreadyExistsError{Message: fmt.Sprintf("a task with the title %v", title)}
	}

	task, err := entities.NewTask(title, description, userEmail)
	if err != nil {
		return err
	}

	err = ct.TaskRepository.CreateTask(*task)
	if err != nil {
		return err
	}

	err = ct.NotificationProvider.SendNotification(*task)
	if err != nil {
		return err
	}
	return nil
}
