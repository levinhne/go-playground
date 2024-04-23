package commands

import (
	"context"

	"github.com/google/wire"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application/ports"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/domain"
)

var CreateTodoHandlerSet = wire.NewSet(NewCreateTodoHandler)

type CreateTodoCommand struct {
	Todo *domain.Todo
}

type CreateTodoHandler struct {
	todoReposiory ports.TodoRepository
}

func NewCreateTodoHandler(todoReposiory ports.TodoRepository) CreateTodoHandler {
	return CreateTodoHandler{
		todoReposiory: todoReposiory,
	}
}

func (h CreateTodoHandler) Handle(ctx context.Context, cmd CreateTodoCommand) error {
	return h.todoReposiory.CreateTodo(ctx, cmd.Todo)
}
