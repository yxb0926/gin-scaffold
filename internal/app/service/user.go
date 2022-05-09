package service

import (
	"context"
	"github.com/google/wire"
	"yxb.com/gin-scaffold/v2/internal/app/dao"
	"yxb.com/gin-scaffold/v2/internal/app/dao/dto"
	"yxb.com/gin-scaffold/v2/internal/app/model"
)

var UserSvcSet = wire.NewSet(wire.Struct(new(UserSvc), "*"))

type UserSvc struct {
	UserRepo *dao.UserRepo
}

func (a *UserSvc) Query(ctx context.Context, req *dto.UserPageReq) (*dto.UserPageRes, error) {
	return a.UserRepo.Query(ctx, req)
}

func (a *UserSvc) Get(ctx context.Context, id int) (*model.User, error) {
	return a.UserRepo.Get(ctx, id)
}

func (a *UserSvc) Create(ctx context.Context, user *model.User) error {
	return a.UserRepo.Create(ctx, user)
}

func (a *UserSvc) Update(ctx context.Context, user *model.User) error {
	return a.UserRepo.Update(ctx, user)
}

func (a *UserSvc) Delete(ctx context.Context, id int) error {
	return a.UserRepo.Delete(ctx, id)
}

func (a *UserSvc) AddRole(ctx context.Context, uid int, rid int) error {
	role := &model.UserRole{UserId: uid, RoleId: rid}
	return a.UserRepo.AddRole(ctx, role)
}
