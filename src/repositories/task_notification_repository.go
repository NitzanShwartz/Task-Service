package repositories

import (
	"github.com/NitzanShwartz/Task-Service/src/entities"
)

type TaskNotificationProvider interface {
	SendNotification(task entities.Task) error
}
