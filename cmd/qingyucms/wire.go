//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	conf2 "github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/db"
	"github.com/iwinder/qingyucms/internal/qycms_blog/server"
	"github.com/iwinder/qingyucms/internal/qycms_blog/service"
)

// wireApp init kratos application.
func wireApp(*conf2.Server, *conf2.Data, *conf.Qycms, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, db.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
