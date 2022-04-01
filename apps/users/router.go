package users

import "github.com/gin-gonic/gin"

func LoadUserRouter(r *gin.Engine) {
	r.GET("users/:userid", FindUser)
	r.GET("users", FindUsers)
	r.POST("users", CreateUser)
	r.DELETE("users/:userid", DeleteUser)
	r.DELETE("users", DeleteUsers)
	r.PATCH("users/:userid", UpdateUser)
}
