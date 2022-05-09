package dao

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"yxb.com/gin-scaffold/v2/internal/app/dao/dto"
	"yxb.com/gin-scaffold/v2/internal/app/model"
)

var UserRepoSet = wire.NewSet(wire.Struct(new(UserRepo), "*"))

type UserRepo struct {
	DB *gorm.DB
}

func (a *UserRepo) Query(ctx context.Context, req *dto.UserPageReq) (*dto.UserPageRes, error) {
	res := &dto.UserPageRes{}
	users := &[]model.User{}
	if err := a.DB.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(users).Error; err != nil {
		return res, err
	}
	var total int64
	if err := a.DB.Model(&model.User{}).Count(&total).Error; err != nil {
		return res, err
	}

	res.List = users
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize
	res.Total = total

	return res, nil
}

func (a *UserRepo) Get(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}
	if err := a.DB.Where("id = ?", id).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (a *UserRepo) Create(ctx context.Context, user *model.User) error {
	return a.DB.Create(user).Error
}

func (a *UserRepo) Update(ctx context.Context, user *model.User) error {
	return a.DB.Where("id = ?", user.ID).Updates(user).Error
}

func (a *UserRepo) Delete(ctx context.Context, id int) error {
	return a.DB.Where("id = ?", id).Delete(&model.User{}).Error
}

func (a *UserRepo) AddRole(ctx context.Context, role *model.UserRole) error {
	return a.DB.Create(role).Error
}
