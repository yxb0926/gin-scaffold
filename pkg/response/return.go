package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(c *gin.Context, code int, msg string) {
	var res response
	res.Success = false
	if msg != "" {
		res.Msg = msg
	}
	res.Code = code
	//c.Set("result", res)
	//c.Set("code", code)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func OK(c *gin.Context, data interface{}) {
	var res response
	res.Success = true
	res.Data = data
	res.Msg = "OK"
	res.Code = http.StatusOK

	// c.Set("result", res)
	// c.Set("code", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {

}
