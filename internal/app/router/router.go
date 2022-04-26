package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	v1 "yxb.com/gin-scaffold/v2/internal/app/api/v1"
)

var Set = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

type Router struct {
	UserApi *v1.UserApi
}

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterApi(app)
	return nil
}

func (a *Router) Prefixes() []string {
	return []string{
		"/api",
	}
}

func (a *Router) RegisterApi(app *gin.Engine) {
	g := app.Group("/api")

	v1 := g.Group("/v1")
	{
		gUser := v1.Group("user")
		{
			gUser.GET(":id", a.UserApi.Get)
			gUser.GET("", a.UserApi.Query)
			gUser.POST("", a.UserApi.Create)
			gUser.DELETE(":id", a.UserApi.Delete)
			gUser.PUT("", a.UserApi.Update)
		}
	}
}
