package adapters

import (
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application/ports"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/domain"
	"github.com/uptrace/bun"
	"golang.org/x/net/context"
)

type TodoRepository struct {
	db *bun.DB
}

func NewTodoRepository(db *bun.DB) ports.TodoRepository {
	return TodoRepository{
		db: db,
	}
}

func (r TodoRepository) GetTodoByID(ctx context.Context, Id string) (*domain.Todo, error) {
	return &domain.Todo{}, nil
}

func (r TodoRepository) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	return nil
}
