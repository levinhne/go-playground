package todo

import (
	"github.com/google/wire"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/adapters"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/handlers"
)

var ProviderSet = wire.NewSet(
	adapters.NewTodoRepository,
	application.NewServiceApplication,
	handlers.NewTodoHandler,
)
