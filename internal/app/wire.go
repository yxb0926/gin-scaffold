//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	v1 "yxb.com/gin-scaffold/v2/internal/app/api/v1"
	"yxb.com/gin-scaffold/v2/internal/app/dao"
	"yxb.com/gin-scaffold/v2/internal/app/router"
	"yxb.com/gin-scaffold/v2/internal/app/service"
)

func BuildInjector() (*Injector, error) {
	panic(wire.Build(
		InitGormDB,
		InitGinEngine,
		v1.ApiSet,
		service.Set,
		dao.Set,
		InjectorSet,
		router.Set,
	))
	return new(Injector), nil
}
