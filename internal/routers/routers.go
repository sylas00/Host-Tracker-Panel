package routers

import (
	"Gin_web/internal/service/hosts"
	"Gin_web/internal/service/users"
	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	// 看Default方法里写的是 New()初始化一个Engine实例 用了日志 和 运行恢复的中间件
	// New()里面有加载路由组 还有重定向相关的一些配置
	r := gin.Default()
	// 挂载路由
	users.LoadUserRouter(r)
	hosts.LoadHostRouter(r)
	return r
}
