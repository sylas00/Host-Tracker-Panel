package users

import (
	"Gin_web/internal/dao"
	"Gin_web/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type jsonUser struct {
	ID       int    `json:"id" gorm:"-"`
	Username string `json:"username" gorm:"-"`
}

var userModel models.User

func FindUser(c *gin.Context) {
	userId := c.Param("userId")
	user := dao.DB.First(&userModel, userId)
	//c.JSON(200, gin.H{"username": user.})
	fmt.Println(user.Name())
	data := user.Scan(&jsonUser{})
	c.JSON(200, data)

}

func FindUsers(c *gin.Context) {

}

// CreateUser 注册用户
func CreateUser(c *gin.Context) {
	//验证json数据
	var user User
	if c.ShouldBindJSON(&user) == nil {
		// 查询用户名是否已经存在
		exist := dao.DB.Where("username = ?", user.Username).First(&userModel)
		if exist == nil {
			//创建用户
			user := models.User{Username: user.Username, Password: user.Password}
			result := dao.DB.Create(&user)
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
	if dao.DB.Delete(&userModel, userId).RowsAffected == 0 {
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
