//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/adapters"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application/ports"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/handlers"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Initialize(db *bun.DB, logger *zap.Logger, grpcServer *grpc.Server) handlers.GrpcHandler {
	panic(wire.Build(
		adapters.TodoRepositorySet,
		wire.Bind(new(ports.TodoRepository), new(adapters.TodoRepository)),
		application.ServiceApplicationSet,
		handlers.GrpcHandlerSet,
	))
}
