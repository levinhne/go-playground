// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/adapters"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/application"
	"github.com/levinhne/grpc-gateway-boilerplate/internal/todo/handlers"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func Initialize(db *bun.DB, logger *zap.Logger) handlers.GrpcHandler {
	todoRepository := adapters.NewTodoRepository(db)
	service := application.NewServiceApplication(todoRepository)
	grpcHandler := handlers.NewGrpcHandler(service, logger)
	return grpcHandler
}
