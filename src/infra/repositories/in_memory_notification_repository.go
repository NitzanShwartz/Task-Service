package repositories

import "github.com/NitzanShwartz/Task-Service/src/entities"

type InMemoryNotificationRepository struct{}

func NewInMemoryNotificationRepository() *InMemoryNotificationRepository {
	return &InMemoryNotificationRepository{}
}

func (im *InMemoryNotificationRepository) SendNotification(task entities.Task) error {
	println("task", task.Title, "was notified about")
	return nil
}
