package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/NitzanShwartz/Task-Service/src/entities"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQNotificationRepository struct {
	connection *amqp.Connection
	queue      amqp.Queue
}

// TODO: test repository

func NewRabbitMQNotificationRepository(connStr string) (*RabbitMQNotificationRepository, error) {
	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"task_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQNotificationRepository{
		connection: conn,
		queue:      q,
	}, nil
}

func (im *RabbitMQNotificationRepository) SendNotification(task entities.Task) error {
	ch, err := im.connection.Channel()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := map[string]string{
		"title":            task.Title,
		"task_description": task.Description,
		"user_email":       task.UserEmail,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil
	}

	err = ch.PublishWithContext(
		ctx,
		"",
		im.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bodyBytes,
		},
	)
	return err
}
