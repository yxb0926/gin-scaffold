package dao

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"yxb.com/gin-scaffold/v2/internal/app/dao/dto"
	"yxb.com/gin-scaffold/v2/internal/app/model"
)

var RoleRepoSet = wire.NewSet(wire.Struct(new(RoleRepo), "*"))

type RoleRepo struct {
	DB *gorm.DB
}

func (a *RoleRepo) Query(ctx context.Context, req *dto.RolePageReq) (*dto.RolePageRes, error) {
	res := &dto.RolePageRes{}
	roles := &[]model.Role{}
	err := a.DB.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(roles).Error
	if err != nil {
		return res, err
	}

	var total int64
	if err = a.DB.Model(&model.Role{}).Count(&total).Error; err != nil {
		return res, err
	}
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize
	res.Total = total
	res.List = roles

	return res, nil
}

func (a *RoleRepo) Get(ctx context.Context, id int) (*model.Role, error) {
	role := &model.Role{}
	err := a.DB.Where("id = ?", id).Find(role).Error

	return role, err
}

func (a *RoleRepo) Create(ctx context.Context, role *model.Role) error {
	return a.DB.Create(role).Error
}

func (a *RoleRepo) Update(ctx context.Context, role *model.Role) error {
	return a.DB.Where("id = ?", role.ID).Updates(role).Error
}

func (a *RoleRepo) Delete(ctx context.Context, id int) error {
	return a.DB.Where("id = ?", id).Delete(&model.Role{}).Error
}

func (a *RoleRepo) RemoveUser(ctx context.Context, role *model.UserRole) error {
	return a.DB.Where("user_id = ? AND role_id = ?", role.UserId, role.RoleId).Delete(&model.UserRole{}).Error
}
