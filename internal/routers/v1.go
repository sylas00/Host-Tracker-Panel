package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func LoadV1(r *gin.Engine) {

	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	// 通过Context的Params方法来获取API参数 :精确匹配  *匹配所有
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})

	r.GET("/hello", func(c *gin.Context) {
		a := c.Query("a")
		b := c.Query("b")
		z := c.DefaultQuery("z", "zz")
		c.String(http.StatusOK, "a， b,z查询参数是: "+a+"  "+b+"  "+z)
	})

	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		asd := c.PostForm("asd")
		//username := c.po("username")
		//password := c.PostForm("password")
		c.String(http.StatusOK, username+" "+password+" "+asd)
	})

	// 创建url组
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
	}
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}