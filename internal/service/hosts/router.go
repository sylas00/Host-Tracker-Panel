package hosts

import "github.com/gin-gonic/gin"

func LoadHostRouter(r *gin.Engine) {
	HostsRouter := r.Group("/hosts")

	HostsRouter.GET("/:hostId", FindHost)
	HostsRouter.GET("", FindHosts)
	HostsRouter.POST("", CreateHost)
	HostsRouter.DELETE("/:hostId", DeleteHost)
	HostsRouter.DELETE("", DeleteHosts)
	HostsRouter.PATCH("/:hostId", UpdateHost)
}
