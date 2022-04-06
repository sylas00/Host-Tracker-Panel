package hosts

import (
	"Gin_web/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindHost(c *gin.Context) {
	userid := c.Param("userid")
	c.String(http.StatusOK, "%s", userid)
}

func FindHosts(c *gin.Context) {

}

// CreateHost 注册用户
func CreateHost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := models.User{Username: username, Password: password}
	// todo create user
	c.String(http.StatusOK, "创建用户%s成功", user.Username)
}

func DeleteHost(c *gin.Context) {

}

func UpdateHost(c *gin.Context) {

}

func DeleteHosts(c *gin.Context) {

}
