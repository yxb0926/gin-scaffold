package service

import (
	"context"
	"github.com/google/wire"
	"yxb.com/gin-scaffold/v2/internal/app/dao"
	"yxb.com/gin-scaffold/v2/internal/app/dao/dto"
	"yxb.com/gin-scaffold/v2/internal/app/model"
)

var RoleSvcSet = wire.NewSet(wire.Struct(new(RoleSvc), "*"))

type RoleSvc struct {
	RoleRepo *dao.RoleRepo
}

func (a *RoleSvc) Query(ctx context.Context, req *dto.RolePageReq) (*dto.RolePageRes, error) {
	return a.RoleRepo.Query(ctx, req)
}

func (a *RoleSvc) Get(ctx context.Context, id int) (*model.Role, error) {
	return a.RoleRepo.Get(ctx, id)
}

func (a *RoleSvc) Create(ctx context.Context, role *model.Role) error {
	return a.RoleRepo.Create(ctx, role)
}

func (a *RoleSvc) Update(ctx context.Context, role *model.Role) error {
	return a.RoleRepo.Update(ctx, role)
}

func (a *RoleSvc) Delete(ctx context.Context, id int) error {
	return a.RoleRepo.Delete(ctx, id)
}

func (a *RoleSvc) RemoveUser(ctx context.Context, role *model.UserRole) error {
	return a.RoleRepo.RemoveUser(ctx, role)
}
