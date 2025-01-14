package services

import "github.com/NitzanShwartz/Task-Service/src/repositories"

type TaskService struct {
	TaskRepository         repositories.TaskRepository
	NotificationRepository repositories.TaskNotificationProvider
}

func NewTaskService(taskRepository repositories.TaskRepository, notificationRepository repositories.TaskNotificationProvider) *TaskService {
	return &TaskService{
		TaskRepository:         taskRepository,
		NotificationRepository: notificationRepository,
	}
}

