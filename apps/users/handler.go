package users

import (
	"Gin_web/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 注册用户
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := models.User{Username: username, Password: password}
	// todo create user
	c.String(http.StatusOK, "创建用户%s成功", user.Username)
}
