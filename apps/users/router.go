package users

import "github.com/gin-gonic/gin"

func LoadUserRouter(r *gin.Engine) {
	r.GET("users/:userId", FindUser)
	r.GET("users", FindUsers)
	r.POST("users", CreateUser)
	r.DELETE("users/:userId", DeleteUser)
	r.DELETE("users", DeleteUsers)
	r.PATCH("users/:userId", UpdateUser)
}
