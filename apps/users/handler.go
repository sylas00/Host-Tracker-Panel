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

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// CreateUser 注册用户
func CreateUser(c *gin.Context) {
	var user User
	if c.BindJSON(&user) == nil {
		exist := models.DB.Where("username = ?", user.Username).First(&models.User{}).Row()
		if exist == nil {
			user := models.User{Username: user.Username, Password: user.Password}
			result := models.DB.Create(&user)
			if result.Error != nil {
				c.JSON(200, gin.H{"status": "失败"})
			} else {
				c.JSON(200, gin.H{"status": "芜湖湖"})
			}
		} else {
			c.JSON(200, gin.H{"status": "用户名已存在"})
		}
	}

}

//todo 写入问题 db

func DeleteUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUsers(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func Logout(c *gin.Context) {

}
