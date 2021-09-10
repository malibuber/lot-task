package routes

import "github.com/gin-gonic/gin"

func GetHallo(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "hadi",
	})
}

func PostHallo(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "hadi",
	})
}
