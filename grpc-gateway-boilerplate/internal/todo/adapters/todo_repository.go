package adapters

import (
	"github.com/google/wire"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application/ports"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/domain"
	"github.com/uptrace/bun"
	"golang.org/x/net/context"
)

var _ ports.TodoRepository = (*TodoRepository)(nil)
var TodoRepositorySet = wire.NewSet(NewTodoRepository)

type TodoRepository struct {
	db *bun.DB
}

func NewTodoRepository(db *bun.DB) TodoRepository {
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
