package controller

import (
	"log"
	"maguas-blog-go/config"
	jwt2 "maguas-blog-go/middleware/jwt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(c *gin.Context) {
	j := &jwt2.JWT{
		[]byte(config.TokenKey),
	}

	claims := jwt2.CustomClaims{
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间一小时
			Issuer:    config.TokenKey,                 // 签名的发现者
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	log.Print(token)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

	return
}

func RefreshToken(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "请求没携带token,无权访问",
		})

		c.Abort()
		return
	}

	var j jwt2.JWT
	refreshToken, err := j.RefreshToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": refreshToken,
	})
	return
}
