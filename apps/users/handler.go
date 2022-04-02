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
	//验证json数据
	var user User
	if c.BindJSON(&user) == nil {
		// 查询用户名是否已经存在
		exist := models.DB.Where("username = ?", user.Username).First(&models.User{})
		if exist == nil {
			//创建用户
			user := models.User{Username: user.Username, Password: user.Password}
			result := models.DB.Create(&user)
			//判断是否创建成功
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

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	if models.DB.Delete(&models.User{}, userId).RowsAffected == 0 {
		c.JSON(200, gin.H{"status": "没有该用户"})
	} else {
		c.JSON(200, gin.H{"status": "删除成功"})
	}
}

func UpdateUser(c *gin.Context) {

}

func DeleteUsers(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func Logout(c *gin.Context) {

}
