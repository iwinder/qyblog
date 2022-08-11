//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/biz"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/conf"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/data/db"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/server"
	"github.com/iwinder/qingyucms/app/qycms_blog/internal/service"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Auth, *conf.Registry, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, db.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
