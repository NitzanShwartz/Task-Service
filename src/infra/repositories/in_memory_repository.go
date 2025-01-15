package repositories

import "github.com/NitzanShwartz/Task-Service/src/entities"

type InMemoryTaskRepository struct {
	repo map[string]entities.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		repo: make(map[string]entities.Task),
	}
}

func (im *InMemoryTaskRepository) DoesTaskExists(taskTitle string) bool {
	_, ok := im.repo[taskTitle]
	return ok
}

func (im *InMemoryTaskRepository) CreateTask(task entities.Task) error {
	im.repo[task.Title] = task
	return nil
}
