package handlers

import (
	"context"

	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application"
	todov1 "github.com/levinhne/grpc-gateway-boilerplate/proto/todo/v1"
	"go.uber.org/zap"
)

type TodoHandler struct {
	app    application.ServiceApplication
	logger *zap.Logger
}

func NewTodoHandler(
	app application.ServiceApplication,
	logger *zap.Logger,
) todov1.TodoServiceServer {
	return &TodoHandler{
		app:    app,
		logger: logger,
	}
}

func (h TodoHandler) CreateTodo(ctx context.Context, in *todov1.CreateTodoRequest) (*todov1.CreateTodoResponse, error) {
	return &todov1.CreateTodoResponse{}, nil
}
