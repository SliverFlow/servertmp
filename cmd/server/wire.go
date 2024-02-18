//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	wire "github.com/google/wire"
	"go.uber.org/zap"
	api "server/api"
	"server/internal/biz"
	"server/internal/config"
	"server/internal/core/server"
	"server/internal/data"
	"server/internal/middleware"
	"server/internal/service"
)

// wireApp init kratos application.
func wireApp(possess *config.Possess, log *zap.Logger) *server.HttpServer {
	panic(wire.Build(api.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, middleware.ProviderSet, server.RunServer))
}
