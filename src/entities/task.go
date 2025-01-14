package entities

import (
	"errors"
)

type Task struct {
	Title       string
	Description string
	UserEmail   string
}

func NewTask(title string, description string, userEmail string) (Task, error) {
	if title == "" || description == "" || userEmail == "" {
		return Task{}, errors.New("title, task and userEmail are mandatory fields")
	}

	return Task{
		Title:       title,
		Description: description,
		UserEmail:   userEmail,
	}, nil
}

func NewEmptyTask() Task {
	return Task{}
}
