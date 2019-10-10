package main

import "github.com/gin-gonic/gin"

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
		"func":    "login",
	})
}

func submit(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
		"func":    "submit",
	})
}

func read(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
		"func":    "read",
	})
}

func main() {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.POST("/login", login)
		v1.POST("/submit", submit)
		v1.POST("/read", read)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
		v2.POST("/read", read)
	}
	r.Run()
}
