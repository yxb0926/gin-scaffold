package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"strconv"
	"yxb.com/gin-scaffold/v2/internal/app/dao/dto"
	"yxb.com/gin-scaffold/v2/internal/app/model"
	"yxb.com/gin-scaffold/v2/internal/app/service"
	"yxb.com/gin-scaffold/v2/pkg/response"
)

var UserApiSet = wire.NewSet(wire.Struct(new(UserApi), "*"))

type UserApi struct {
	UserSvc *service.UserSvc
}

// Query 分页查询用户信息
func (a *UserApi) Query(c *gin.Context) {
	req := &dto.UserPageReq{}
	if err := c.BindJSON(req); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	data, err := a.UserSvc.Query(c.Request.Context(), req)
	if err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, data)
}

// Get 按照ID查询用户信息
func (a *UserApi) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	data, err := a.UserSvc.Get(c.Request.Context(), id)
	if err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, data)
}

// Create 新建用户
func (a *UserApi) Create(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(&user); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}

	if err := a.UserSvc.Create(c.Request.Context(), user); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, user)
}

// Delete 按照ID删除用户
func (a *UserApi) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	if err = a.UserSvc.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, id)
}

// Update 升级用户信息
func (a *UserApi) Update(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	if err := a.UserSvc.Update(c.Request.Context(), user); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, user)
}

// AddRole 给用户增加角色
// POST: api/v1/user/:id/role/:rid
func (a *UserApi) AddRole(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	rid, err := strconv.Atoi(c.Param("rid"))
	if err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	if err = a.UserSvc.AddRole(c.Request.Context(), uid, rid); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, nil)
}
