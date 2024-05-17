package router

import (
	"github.com/gin-gonic/gin"
	v1 "wxapp/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	if gin.Mode() == gin.DebugMode {
		gin.SetMode(gin.DebugMode)
	}
	r.Use(gin.Logger())
	r.GET("/ping", v1.HelloWorld())
	return r
}
