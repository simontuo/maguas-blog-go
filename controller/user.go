package controller

import (
	"maguas-blog-go/database"
	"maguas-blog-go/model"
	"maguas-blog-go/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSearch(c *gin.Context) {
	db, _ := database.Connect()
	defer db.Close()

	var users []model.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func UserShow(c *gin.Context) {
	db, _ := database.Connect()
	defer db.Close()

	var user model.User
	db.Where("id = ?", c.Param("user")).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UserUpdate(c *gin.Context) {
	db, _ := database.Connect()
	defer db.Close()

	var user model.User
	row := db.Model(&user).Where("id = ?", c.Param("user")).Updates(model.User{
		Name:   c.PostForm("name"),
		Phone:  c.PostForm("phone"),
		Email:  c.PostForm("email"),
		Avatar: c.PostForm("avatar"),
		Password: c.PostForm("password"),
	}).RowsAffected

	if row < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "update fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "update success",
	})
}

func UserCreate(c *gin.Context) {
	var postData validation.User
	if err := c.ShouldBind(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	db, _ := database.Connect()
	defer db.Close()

	user := model.User{
		Name:     c.PostForm("name"),
		Phone:    c.PostForm("phone"),
		Email:    c.PostForm("email"),
		Avatar:   c.PostForm("avatar"),
		Password: c.PostForm("password"),
	}

	if db.NewRecord(user) {
		db.Create(&user)

		if db.NewRecord(user) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "create fail",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "create success",
		})
		return
	}

	c.JSON(http.StatusForbidden, gin.H{
		"msg": "this user already exist",
	})
}

func UserDelete(c *gin.Context) {
	db, _ := database.Connect()
	defer db.Close()

	var user model.User
	row := db.Where("id = ?", c.Param("user")).Delete(&user).RowsAffected

	if row < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "delete fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "delete success",
	})
}