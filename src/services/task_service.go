package services

import (
	"github.com/NitzanShwartz/Task-Service/src/repositories"
	usecases "github.com/NitzanShwartz/Task-Service/src/use_cases"
)

type TaskService struct {
	CreateTaskUseCase usecases.CreateTask
	NotifyTaskUseCase usecases.NotifyTask
}

func NewTaskService(taskRepository repositories.TaskRepository, notificationRepository repositories.TaskNotificationProvider) *TaskService {
	return &TaskService{
		CreateTaskUseCase: usecases.CreateTask{
			TaskRepository: taskRepository,
		},
		NotifyTaskUseCase: usecases.NotifyTask{
			NotificationRepository: notificationRepository,
		},
	}
}

func (ts *TaskService) CreateTask(title string, description string, userEmail string) error {
	task, err := ts.CreateTaskUseCase.Execute(title, description, userEmail)
	if err != nil {
		return err
	}

	err = ts.NotifyTaskUseCase.Execute(task)
	return err
}
