package routers

import "github.com/gin-gonic/gin"

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	// 创建路由
	r := gin.Default()
	LoadV1(r)
	return r
}
