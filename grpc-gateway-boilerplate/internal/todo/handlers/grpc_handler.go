package handlers

import (
	"context"

	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application"
	todov1 "github.com/levinhne/grpc-gateway-boilerplate/proto/todo/v1"
	"go.uber.org/zap"
)

type GrpcHandler struct {
	app    application.ServiceApplication
	logger *zap.Logger
}

func NewGrpcHandler(app application.ServiceApplication, logger *zap.Logger) GrpcHandler {
	return GrpcHandler{
		app:    app,
		logger: logger,
	}
}

func (h GrpcHandler) CreateTodo(ctx context.Context, in *todov1.CreateTodoRequest) (*todov1.CreateTodoResponse, error) {
	return &todov1.CreateTodoResponse{}, nil
}
