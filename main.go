package main

import (
	"lot-task/routes"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	routers := gin.Default()
	/* 	routers.Use(check()) */
	v1 := routers.Group("/hallo")
	{
		v1.GET("/", authMiddleWare(), routes.GetHallo)
		v1.POST("/", authMiddleWare(), routes.PostHallo)
	}
	routers.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "path fail"})
	})
	routers.Run(":3000")
}

/*
func check() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path != "/hallo" {
			c.AbortWithStatusJSON(500, gin.H{"message": "path fail"})
		}
	}
} */

func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(401, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		if token == "qwerty" {
			c.Next()
		} else {
			c.JSON(401, gin.H{"error": "Not authenticated."})
			c.Abort()
		}
	}
}
