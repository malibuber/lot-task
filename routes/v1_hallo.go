package routes

import "github.com/gin-gonic/gin"

func GetHallo(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "get",
	})
}

func PostHallo(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "post",
	})
}

func Test(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "post",
	})
}
