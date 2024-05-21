package handlers

import (
	"context"

	"github.com/google/wire"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application"
	todov1 "github.com/levinhne/grpc-gateway-boilerplate/proto/todo/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var GrpcHandlerSet = wire.NewSet(NewGrpcHandler)

type GrpcHandler struct {
	app    application.ServiceApplication
	logger *zap.Logger
}

func NewGrpcHandler(
	grpcServer *grpc.Server,
	app application.ServiceApplication,
	logger *zap.Logger,
) todov1.TodoServiceServer {
	srv := GrpcHandler{
		app:    app,
		logger: logger,
	}
	todov1.RegisterTodoServiceServer(grpcServer, &srv)
	return srv
}

func (h GrpcHandler) CreateTodo(ctx context.Context, in *todov1.CreateTodoRequest) (*todov1.CreateTodoResponse, error) {
	return &todov1.CreateTodoResponse{}, nil
}
