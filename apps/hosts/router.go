package hosts

import "github.com/gin-gonic/gin"

func LoadHostRouter(r *gin.Engine) {
	r.GET("hosts/:hostId", FindHost)
	r.GET("hosts", FindHosts)
	r.POST("hosts", CreateHost)
	r.DELETE("hosts/:hostId", DeleteHost)
	r.DELETE("hosts", DeleteHosts)
	r.PATCH("hosts/:hostId", UpdateHost)
}
