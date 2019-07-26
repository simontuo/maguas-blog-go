package tokencontroller

import (
	"log"
	"maguas-blog-go/config"
	"maguas-blog-go/database"
	jwt2 "maguas-blog-go/middleware/jwt"
	"maguas-blog-go/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(c *gin.Context) {
	j := &jwt2.JWT{
		[]byte(config.TokenKey),
	}

	db, _ := database.Connect()
	defer db.Close()
	var user model.User
	db.First(&user)

	claims := jwt2.CustomClaims{
		user.ID,
		user.Name,
		user.Phone,
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

}
