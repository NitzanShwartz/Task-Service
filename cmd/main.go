package main

import (
	grpcService "github.com/NitzanShwartz/Task-Service/src/infra/api/grpcServer"
	"github.com/NitzanShwartz/Task-Service/src/infra/repositories"
)

func main() {
	inMemoryRepo := repositories.NewInMemoryTaskRepository()
	inMemoryNotificationRepo := repositories.NewInMemoryNotificationRepository()

	grpcServer := grpcService.NewGRPCServer(inMemoryRepo, inMemoryNotificationRepo)
	grpcServer.Serve(5052)
}
