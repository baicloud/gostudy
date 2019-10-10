package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("moreXML", func(c *gin.Context) {
		type MessageRecord struct {
			Name    string
			Message string
			Number  int
		}
		var msg MessageRecord
		msg.Name = "lena"
		msg.Message = "key"
		msg.Number = 123
		c.XML(http.StatusOK, msg)
	})
	r.Run()
}
