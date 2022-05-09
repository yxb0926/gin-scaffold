package dao

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"yxb.com/gin-scaffold/v2/internal/app/dao/dto"
	"yxb.com/gin-scaffold/v2/internal/app/model"
)

var MenuRepoSet = wire.NewSet(wire.Struct(new(MenuRepo), "*"))

type MenuRepo struct {
	DB *gorm.DB
}

func (a *MenuRepo) Query(ctx context.Context, req *dto.MenuPageReq) (*dto.MenuPageRes, error) {
	res := &dto.MenuPageRes{}
	menus := &[]model.Menu{}
	err := a.DB.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(menus).Error
	if err != nil {
		return res, err
	}

	var total int64
	err = a.DB.Model(&model.Menu{}).Count(&total).Error
	if err != nil {
		return res, err
	}

	res.Total = total
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize
	res.List = menus

	return res, nil
}

func (a *MenuRepo) Get(ctx context.Context, id int) (*model.Menu, error) {
	menu := &model.Menu{}
	err := a.DB.Where("id = ?", id).Find(menu).Error
	return menu, err
}

func (a *MenuRepo) Create(ctx context.Context, menu *model.Menu) error {
	return a.DB.Create(menu).Error
}

func (a *MenuRepo) Update(ctx context.Context, menu *model.Menu) error {
	return a.DB.Where("id = ?", menu.ID).Updates(menu).Error
}

func (a *MenuRepo) Delete(ctx context.Context, id int) error {
	return a.DB.Where("id = ?", id).Delete(&model.Menu{}).Error
}
