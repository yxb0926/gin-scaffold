package service

import (
	"context"
	"github.com/google/wire"
	"yxb.com/gin-scaffold/v2/internal/app/dao"
	"yxb.com/gin-scaffold/v2/internal/app/dao/dto"
	"yxb.com/gin-scaffold/v2/internal/app/model"
)

var MenuSvcSet = wire.NewSet(wire.Struct(new(MenuSvc), "*"))

type MenuSvc struct {
	MenuRepo *dao.MenuRepo
}

func (a *MenuSvc) Query(ctx context.Context, req *dto.MenuPageReq) (*dto.MenuPageRes, error) {
	return a.MenuRepo.Query(ctx, req)
}

func (a *MenuSvc) Get(ctx context.Context, id int) (*model.Menu, error) {
	return a.MenuRepo.Get(ctx, id)
}

func (a *MenuSvc) Create(ctx context.Context, menu *model.Menu) error {
	return a.MenuRepo.Create(ctx, menu)
}

func (a *MenuSvc) Update(ctx context.Context, menu *model.Menu) error {
	return a.MenuRepo.Update(ctx, menu)
}

func (a *MenuSvc) Delete(ctx context.Context, id int) error {
	return a.MenuRepo.Delete(ctx, id)
}
