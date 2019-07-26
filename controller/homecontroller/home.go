package homecontroller

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is a index page",
	})
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is a test page",
	})
}
