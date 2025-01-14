package usecases

import (
	"github.com/NitzanShwartz/Task-Service/src/entities"
	"github.com/NitzanShwartz/Task-Service/src/repositories"
)

type NotifyTask struct {
	NotificationRepository repositories.TaskNotificationProvider
}

func NewNotifyTask(notificationRepository repositories.TaskNotificationProvider) *NotifyTask {
	return &NotifyTask{
		NotificationRepository: notificationRepository,
	}
}

func (nt *NotifyTask) Execute(task entities.Task) error {
	err := nt.NotificationRepository.SendNotification(task)
	return err
}
