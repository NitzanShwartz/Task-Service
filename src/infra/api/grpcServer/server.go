package grpc_impl

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/NitzanShwartz/Task-Service/src/repositories"
	"github.com/NitzanShwartz/Task-Service/src/services"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	TaskServiceServer
	TaskService *services.TaskService
}

func NewGRPCServer(taskRepo repositories.TaskRepository, notificationRepo repositories.TaskNotificationProvider) *GRPCServer {
	return &GRPCServer{
		TaskService: services.NewTaskService(taskRepo, notificationRepo),
	}
}

func (g *GRPCServer) CreateTask(ctx context.Context, taskMessage *TaskMessage) (*EmptyResponse, error) {
	return &EmptyResponse{}, g.TaskService.CreateTask(taskMessage.TaskTitle, taskMessage.Task, taskMessage.UserEmail)
}

func (g *GRPCServer) Serve(port int) {
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		os.Exit(1) // TODO: log.Fatalf instead of this
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterTaskServiceServer(grpcServer, g)
	grpcServer.Serve(list)
}
