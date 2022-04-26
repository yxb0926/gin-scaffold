package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"strconv"
	"yxb.com/gin-scaffold/v2/internal/app/model"
	"yxb.com/gin-scaffold/v2/internal/app/service"
	"yxb.com/gin-scaffold/v2/pkg/response"
)

var UserApiSet = wire.NewSet(wire.Struct(new(UserApi), "*"))

type UserApi struct {
	UserSvc *service.UserSvc
}

func (a *UserApi) Query(c *gin.Context) {
	req := &model.UserPagReq{}
	if err := c.BindJSON(req); err != nil {
		response.Error(c, response.ErrInvalidParentCode, err.Error())
		return
	}
	data, err := a.UserSvc.Query(c.Request.Context(), req)
	if err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, data)
}

func (a *UserApi) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, response.ErrInvalidParentCode, err.Error())
		return
	}
	data, err := a.UserSvc.Get(c.Request.Context(), id)
	if err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, data)
}

func (a *UserApi) Create(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(&user); err != nil {
		response.Error(c, response.ErrInvalidParentCode, err.Error())
		return
	}

	if err := a.UserSvc.Create(c.Request.Context(), user); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, user)
}

func (a *UserApi) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, response.ErrInvalidParentCode, err.Error())
		return
	}
	if err = a.UserSvc.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, id)
}

func (a *UserApi) Update(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		response.Error(c, response.ErrInvalidParentCode, err.Error())
		return
	}
	if err := a.UserSvc.Update(c.Request.Context(), user); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, user)
}
