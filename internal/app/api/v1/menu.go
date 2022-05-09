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

var MenuApiSet = wire.NewSet(wire.Struct(new(MenuApi), "*"))

type MenuApi struct {
	MenuSvc *service.MenuSvc
}

func (a *MenuApi) Query(c *gin.Context) {
	req := &dto.MenuPageReq{}
	if err := c.BindJSON(req); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}

	data, err := a.MenuSvc.Query(c.Request.Context(), req)
	if err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, data)
}
func (a *MenuApi) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}

	data, err := a.MenuSvc.Get(c.Request.Context(), id)
	if err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, data)
}
func (a *MenuApi) Create(c *gin.Context) {
	menu := &model.Menu{}
	if err := c.BindJSON(menu); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	if err := a.MenuSvc.Create(c.Request.Context(), menu); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, menu)
}
func (a *MenuApi) Update(c *gin.Context) {
	menu := &model.Menu{}
	if err := c.BindJSON(menu); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	if err := a.MenuSvc.Update(c.Request.Context(), menu); err != nil {
		response.Error(c, response.ErrSvcCode, err.Error())
		return
	}
	response.OK(c, menu)
}
func (a *MenuApi) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	if err = a.MenuSvc.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, response.ErrArgsCode, err.Error())
		return
	}
	response.OK(c, id)
}
