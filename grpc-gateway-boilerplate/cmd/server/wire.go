//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo"
	todov1 "github.com/levinhne/grpc-gateway-boilerplate/proto/todo/v1"
	"google.golang.org/grpc"
)

type App struct{}

func New(server *grpc.Server, todoSrv todov1.TodoServiceServer) *App {
	todov1.RegisterTodoServiceServer(server, todoSrv)
	return &App{}
}

func Initialize(server *grpc.Server) (*App, error) {
	panic(wire.Build(
		New,
		todo.ProviderSet,
	))
}
