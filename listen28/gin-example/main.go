package main

import (
	"github.com/gin-gonic/gin"
)

func testHandle(c *gin.Context) {
	c.Request.Cookie("user_cook")
	c.JSON(200, gin.H{
		"message": "test",
	})
}

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", testHandle)
	r.Run(":8080")
}
