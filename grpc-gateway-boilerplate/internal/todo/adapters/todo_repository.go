package adapters

import (
	"fmt"
	"sync"

	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application/ports"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/domain"
	"golang.org/x/net/context"
)

type TodoRepository struct {
	todos map[string]*domain.Todo
	lock  *sync.RWMutex
}

func NewTodoRepository() ports.TodoRepository {
	return TodoRepository{
		todos: map[string]*domain.Todo{},
		lock:  &sync.RWMutex{},
	}
}

func (r TodoRepository) GetTodoByID(_ context.Context, ID string) (*domain.Todo, error) {
	r.lock.RLock()
	todo, ok := r.todos[ID]
	if !ok {
		return nil, fmt.Errorf("todo not found")
	}
	defer r.lock.RUnlock()
	return todo, nil
}

func (r TodoRepository) CreateTodo(_ context.Context, todo *domain.Todo) error {
	r.lock.Lock()
	r.todos[todo.ID] = todo
	defer r.lock.Unlock()
	return nil
}
