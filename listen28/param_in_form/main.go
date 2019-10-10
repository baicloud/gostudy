package main

import "github.com/gin-gonic/gin"

func main() {
	//Default返回默认路由引擎
	r:=gin.Default()
	r.POST("/user/search",func(c *gin.Context){
		username:=c.DefaultPostForm("username","这个")
		address:=c.PostForm("address")
		//输出json给调用方
		c.JSON(200,gin.H{
			"message":"pong",
			"username":username,
			"address":address,
		})
	})
	r.Run()
}
