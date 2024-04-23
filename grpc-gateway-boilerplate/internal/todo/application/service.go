package application

import (
	"context"

	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application/commands"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application/ports"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application/queries"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/domain"
)

var _ ServiceApplication = (*Service)(nil)

type ServiceApplication interface {
	GetTodo(ctx context.Context, query queries.GetTodo) (*domain.Todo, error)
	CreateTodo(ctx context.Context, cmd commands.CreateTodoCommand) error
}

type Service struct {
	commands Commands
	queries  Queries
}

type Queries struct {
	GetTodo queries.GetTodoHandler
}

type Commands struct {
	CreateTodo commands.CreateTodoHandler
}

func NewServiceApplication(
	todoRepository ports.TodoRepository,
) Service {
	return Service{
		commands: Commands{
			CreateTodo: commands.NewCreateTodoHandler(todoRepository),
		},
		queries: Queries{
			GetTodo: queries.NewTodoHandler(todoRepository),
		},
	}
}

func (s Service) GetTodo(ctx context.Context, query queries.GetTodo) (*domain.Todo, error) {
	return s.queries.GetTodo.Handle(ctx, query)
}

func (s Service) CreateTodo(ctx context.Context, cmd commands.CreateTodoCommand) error {
	return s.commands.CreateTodo.Handle(ctx, cmd)
}
