package hosts

import "github.com/gin-gonic/gin"

func LoadHostRouter(r *gin.Engine) {
	r.GET("hosts/:host_id", FindHost)
	r.GET("hosts", FindHosts)
	r.POST("hosts", CreateHost)
	r.DELETE("hosts/:host_id", DeleteHost)
	r.DELETE("hosts", DeleteHosts)
	r.PATCH("hosts/:host_id", UpdateHost)
}
