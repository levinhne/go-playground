package queries

import (
	"context"
	"fmt"

	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application/ports"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/domain"
)

type GetTodo any

type GetTodoByID struct {
	ID string
}

type GetTodoHandler struct {
	todoRepository ports.TodoRepository
}

func NewTodoHandler(todoRepository ports.TodoRepository) GetTodoHandler {
	return GetTodoHandler{
		todoRepository: todoRepository,
	}
}

func (h GetTodoHandler) Handle(ctx context.Context, query GetTodo) (*domain.Todo, error) {
	switch q := query.(type) {
	case GetTodoByID:
		return h.todoRepository.GetTodoByID(ctx, q.ID)
	}
	return nil, fmt.Errorf("query not found")
}
