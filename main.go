package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
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

	//查询参数
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
