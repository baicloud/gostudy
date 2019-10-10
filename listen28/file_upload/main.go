package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	r:=gin.Default()
	r.POST("/upload",func(c *gin.Context){
		file,err:=c.FormFile("file")
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"message":err.Error(),
			})
			return
		}
		log.Println(file)
		log.Println(file.Filename)
		dst:=fmt.Sprintf("/tmp/%s",file.Filename)
		c.SaveUploadedFile(file,dst)
		c.JSON(http.StatusOK,gin.H{
			"message":fmt.Sprintf("'%s' uploaded!",file.Filename),
		})
	})
	r.Run()

}