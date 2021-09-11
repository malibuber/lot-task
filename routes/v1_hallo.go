package routes

//  "request":{"name":"Vitali","company":"LotInternet","task":"make_cofee"}}
import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type HalloOrder struct {
	Name    *string `json:"name" form:"name"`
	Company *string `json:"company" form:"company"`
	Task    *string `json:"task" form:"task"`
}

func GetHallo(c *gin.Context) {

	/* 	name := c.Query("name") */
	var h *HalloOrder
	if err := c.ShouldBindQuery(&h); err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "something went wrong",
		})
	}
	c.JSON(200, gin.H{
		"result":  "Hallo," + *h.Name + " from " + *h.Company + " Please do task: " + *h.Task,
		"request": *h,
	})
}

func PostHallo(c *gin.Context) {

	var h *HalloOrder

	if err := c.ShouldBindBodyWith(&h, binding.JSON); err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "something went wrong",
		})
	}
	fmt.Println(*h)
	c.JSON(200, gin.H{
		"result":  "Hallo," + *h.Name + " from " + *h.Company + " Please do task: " + *h.Task,
		"request": *h,
	})
}
