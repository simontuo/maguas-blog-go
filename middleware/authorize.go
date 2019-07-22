package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 身份验证
func Auth() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "this is a middleware page",
		})
		return
	}
}