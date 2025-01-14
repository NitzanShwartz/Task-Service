package repositories

import (
	"github.com/NitzanShwartz/Task-Service/src/entities"
)

type TaskRepository interface {
	DoesTaskExists(title string) bool
	CreateTask(task entities.Task) error
}
