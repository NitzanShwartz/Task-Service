package main

import (
	grpcService "github.com/NitzanShwartz/Task-Service/src/infra/api/grpcServer"
	"github.com/NitzanShwartz/Task-Service/src/infra/repositories"
)

const port = 5052

func main() {
	inMemoryRepo := repositories.NewInMemoryTaskRepository()
	rabbitmqRepository, err := repositories.NewRabbitMQNotificationRepository("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	grpcServer := grpcService.NewGRPCServer(inMemoryRepo, rabbitmqRepository)
	grpcServer.Serve(port)
}
