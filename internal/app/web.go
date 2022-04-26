package app

import (
	"github.com/gin-gonic/gin"
	"yxb.com/gin-scaffold/v2/internal/app/router"
)

func InitGinEngine(r router.IRouter) *gin.Engine {
	app := gin.New()
	r.Register(app)

	return app
}
