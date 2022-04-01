package users

import "github.com/gin-gonic/gin"

func LoadUserRouter(r *gin.Engine) {
	r.POST("register", UserRegister)
	r.GET("info/:userid", UserInfo)
}
