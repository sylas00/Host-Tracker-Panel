package routers

import (
	"Gin_web/apps/users"
	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	// 创建路由
	r := gin.Default()
	LoadV1(r)
	users.LoadUserRouter(r)
	return r
}
