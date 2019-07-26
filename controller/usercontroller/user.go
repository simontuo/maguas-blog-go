package usercontroller

import (
	"maguas-blog-go/database"
	"maguas-blog-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	db, _ := database.Connect()
	defer db.Close()

	var users []model.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func Show(c *gin.Context) {
	db, _ := database.Connect()
	defer db.Close()

	var user model.User
	db.Debug().Where("id = ?", c.Param("user")).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Update(c *gin.Context) {
	db, _ := database.Connect()
	defer db.Close()

	var user model.User
	row := db.Model(&user).Where("id = ?", c.Param("user")).Updates(model.User{
		Name:   c.PostForm("name"),
		Phone:  c.PostForm("phone"),
		Email:  c.PostForm("email"),
		Avatar: c.PostForm("avatar"),
	}).RowsAffected

	if row > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "update success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "update fail",
		})
	}

}

func Create(c *gin.Context) {
	db, _ := database.Connect()
	defer db.Close()

	user := model.User{
		Name:     c.PostForm("name"),
		Phone:    c.PostForm("phone"),
		Email:    c.PostForm("email"),
		Avatar:   c.PostForm("avatar"),
		Password: c.PostForm("password"),
	}

	err := user.Verify()
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": err,
		})
	}

	if db.NewRecord(user) {
		db.Create(&user)

		if db.NewRecord(user) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "create fail",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "create success",
			})
		}

	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "该用户数据已存在。",
		})
	}
}

func Delete(c *gin.Context) {
	db, _ := database.Connect()
	defer db.Close()

	var user model.User
	row := db.Where("id = ?", c.Param("user")).Delete(&user).RowsAffected

	if row > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete fail",
		})
	}
}
