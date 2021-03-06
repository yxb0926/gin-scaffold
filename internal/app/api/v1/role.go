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

var RoleApiSet = wire.NewSet(wire.Struct(new(RoleApi), "*"))

type RoleApi struct {
	RoleSvc *service.RoleSvc
}

func (a *RoleApi) Query(c *gin.Context) {
	req := &dto.RolePageReq{}
	if err := c.BindJSON(req); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	data, err := a.RoleSvc.Query(c.Request.Context(), req)
	if err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, data)
}

func (a *RoleApi) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	data, err := a.RoleSvc.Get(c.Request.Context(), id)
	if err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, data)
}

func (a *RoleApi) Create(c *gin.Context) {
	role := model.Role{}
	if err := c.BindJSON(&role); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}

	if err := a.RoleSvc.Create(c.Request.Context(), &role); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	response.OK(c, role)
}

func (a *RoleApi) Update(c *gin.Context) {
	role := &model.Role{}
	if err := c.BindJSON(role); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	if err := a.RoleSvc.Update(c.Request.Context(), role); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, role)
}

func (a *RoleApi) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	if err = a.RoleSvc.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, id)
}

// RemoveUser ???role???????????????
// Delete: api/v1/role/:id/user/:uid
func (a *RoleApi) RemoveUser(c *gin.Context) {
	role := &model.UserRole{}
	if err := c.BindJSON(role); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	if err := a.RoleSvc.RemoveUser(c.Request.Context(), role); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, role)
}

// AddMenu ???role??????menu
// Post: api/v1/role/:id/menu/:mid
func (a *RoleApi) AddMenu(c *gin.Context) {

}

// RemoveMenu ???role??????menu
// Delete: api/v1/role/:id/menu/:mid
func (a *RoleApi) RemoveMenu(c *gin.Context) {

}
