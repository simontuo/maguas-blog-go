package usercontroller

import (
	"maguas-blog-go/database"
	"maguas-blog-go/model"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	db, _ := database.Connect()
	defer db.Close()

	var user model.User
	db.Preload("Articles.Comments").First(&user)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func Show(c *gin.Context) {

}

func Update(c *gin.Context) {

}
