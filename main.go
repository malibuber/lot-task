package main

import "github.com/gin-gonic/gin"

func main() {
	routers := gin.Default()
	routers.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "test",
		})
	})
	routers.Run()
}
