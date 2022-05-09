// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"yxb.com/gin-scaffold/v2/internal/app/api/v1"
	"yxb.com/gin-scaffold/v2/internal/app/dao"
	"yxb.com/gin-scaffold/v2/internal/app/router"
	"yxb.com/gin-scaffold/v2/internal/app/service"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, error) {
	db, err := InitGormDB()
	if err != nil {
		return nil, err
	}
	userRepo := &dao.UserRepo{
		DB: db,
	}
	userSvc := &service.UserSvc{
		UserRepo: userRepo,
	}
	userApi := &v1.UserApi{
		UserSvc: userSvc,
	}
	roleRepo := &dao.RoleRepo{
		DB: db,
	}
	roleSvc := &service.RoleSvc{
		RoleRepo: roleRepo,
	}
	roleApi := &v1.RoleApi{
		RoleSvc: roleSvc,
	}
	routerRouter := &router.Router{
		UserApi: userApi,
		RoleApi: roleApi,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine: engine,
	}
	return injector, nil
}
