package ports

import (
	"context"

	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/domain"
)

type TodoRepository interface {
	GetTodoByID(ctx context.Context, ID string) (*domain.Todo, error)
	ListTodos(ctx context.Context) ([]*domain.Todo, error)
	CreateTodo(ctx context.Context, todo *domain.Todo) error
}
