package routes

//  "request":{"name":"Vitali","company":"LotInternet","task":"make_cofee"}}
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type HalloOrder struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Task    string `json:"task"`
}

func GetHallo(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "get",
	})
}

func PostHallo(c *gin.Context) {

	var h HalloOrder
	if err := c.ShouldBindBodyWith(&h, binding.JSON); err != nil {
		c.JSON(404, gin.H{
			"message": "something went wrong",
		})
	}
	c.JSON(200, gin.H{
		"result": "Hallo," + h.Name + " from " + h.Company + " Please do task: " + h.Task,
	})
}
