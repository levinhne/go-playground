package handlers

import (
	"context"
	"log"

	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application"
	todov1 "github.com/levinhne/grpc-gateway-boilerplate/proto/todo/v1"
)

type TodoHandler struct {
	app application.ServiceApplication
}

func NewTodoHandler(
	app application.ServiceApplication,
) todov1.TodoServiceServer {
	return &TodoHandler{
		app: app,
	}
}

func (h TodoHandler) GetTodo(ctx context.Context, in *todov1.GetTodoRequest) (*todov1.GetTodoResponse, error) {
	log.Println(in.GetName())
	return nil, nil
}

func (h TodoHandler) ListTodos(ctx context.Context, in *todov1.ListTodosRequest) (*todov1.ListTodosResponse, error) {
	return nil, nil
}

func (h TodoHandler) CreateTodo(ctx context.Context, in *todov1.CreateTodoRequest) (*todov1.CreateTodoResponse, error) {
	return &todov1.CreateTodoResponse{}, nil
}

func (h TodoHandler) ListTodoTags(ctx context.Context, in *todov1.ListTodoTagsRequest) (*todov1.ListTodoTagsResponse, error) {
	log.Println(in.Parent)
	return nil, nil
}
