package routes

//  "request":{"name":"Vitali","company":"LotInternet","task":"make_cofee"}}
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type HalloOrder struct {
	Name    *string `json:"name" form:"name" binding:"required"`
	Company *string `json:"company" form:"company" binding:"required"`
	Task    string  `json:"task" form:"task"`
}

func GetHallo(c *gin.Context) {

	/* 	name := c.Query("name") */
	var h *HalloOrder
	if err := c.ShouldBindQuery(&h); err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "missing value",
		})
		return
	}
	if h.Task == "" {
		h.Task = "make_coffe"
	}
	c.JSON(200, gin.H{
		"result":  "Hallo," + *h.Name + " from " + *h.Company + " Please do task: " + h.Task,
		"request": *h,
	})
}

func PostHallo(c *gin.Context) {

	var h *HalloOrder

	if err := c.ShouldBindBodyWith(&h, binding.JSON); err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "missing value",
		})
		return
	}
	if h.Task == "" {
		h.Task = "make_coffe"
	}
	c.JSON(200, gin.H{
		"result":  "Hallo," + *h.Name + " from " + *h.Company + " Please do task: " + h.Task,
		"request": *h,
	})
}
