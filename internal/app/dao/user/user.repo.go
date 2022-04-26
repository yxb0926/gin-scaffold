package user

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"yxb.com/gin-scaffold/v2/internal/app/model"
)

var RepoSet = wire.NewSet(wire.Struct(new(Repo), "*"))

type Repo struct {
	DB *gorm.DB
}

func (a *Repo) Query(ctx context.Context, req *model.UserPagReq) (*model.UserPageRes, error) {
	user := &[]model.User{}
	if err := a.DB.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(user).Error; err != nil {
		return nil, err
	}
	var total int64
	if err := a.DB.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, err
	}

	res := &model.UserPageRes{}
	res.List = user
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize
	res.Total = total

	return res, nil
}

func (a *Repo) Get(ctx context.Context, id int) (*model.User, error) {
	var user *model.User
	if err := a.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (a *Repo) Create(ctx context.Context, user *model.User) error {
	return a.DB.Create(user).Error
}

func (a *Repo) Update(ctx context.Context, user *model.User) error {
	return a.DB.Updates(user).Error
}

func (a *Repo) Delete(ctx context.Context, id int) error {
	return a.DB.Where("id = ?", id).Delete(&model.User{}).Error
}
