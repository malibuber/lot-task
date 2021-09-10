package main

import (
	"lot-task/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	routers := gin.Default()

	v1 := routers.Group("hallo")
	{
		v1.GET("/", routes.GetHallo)
		v1.GET("/", routes.PostHallo)
	}
	routers.Run(":3000")
}
