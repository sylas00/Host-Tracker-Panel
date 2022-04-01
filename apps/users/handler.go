package users

import (
	"Gin_web/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindUser(c *gin.Context) {
	userid := c.Param("userid")
	c.String(http.StatusOK, "%s", userid)
}

func FindUsers(c *gin.Context) {

}

// CreateUser 注册用户
func CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := models.User{Username: username, Password: password}
	// todo create user
	c.String(http.StatusOK, "创建用户%s成功", user.Username)
}

func DeleteUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUsers(c *gin.Context) {

}
