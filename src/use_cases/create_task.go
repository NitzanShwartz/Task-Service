package usecases

import (
	"errors"
	"fmt"
	"net/mail"

	"github.com/NitzanShwartz/Task-Service/src/entities"
	"github.com/NitzanShwartz/Task-Service/src/repositories"
	"github.com/NitzanShwartz/Task-Service/src/use_cases/exceptions"
)

type CreateTask struct {
	TaskRepository repositories.TaskRepository
}

func NewCreateTask(taskRepsitory repositories.TaskRepository) *CreateTask {
	return &CreateTask{
		TaskRepository: taskRepsitory,
	}
}

func (ct *CreateTask) Execute(title string, description string, userEmail string) error {
	if title == "" || description == "" || userEmail == "" {
		return errors.New("title, task and userEmail are mandatory fields")
	}

	if _, err := mail.ParseAddress(userEmail); err != nil {
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

	return nil
}
