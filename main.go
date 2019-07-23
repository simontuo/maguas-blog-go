package main

import (
	"github.com/gin-gonic/gin"
	"maguas-blog-go/controller/articlecontroller"
	"maguas-blog-go/controller/commentcontroller"
	"maguas-blog-go/controller/homecontroller"
	"maguas-blog-go/controller/tagcontroller"
	"maguas-blog-go/controller/usercontroller"
	"maguas-blog-go/middleware"
)

func main() {
	r := gin.Default()
	// 非验证接口
	r.GET("/", homecontroller.Index)
	r.GET("/articles", articlecontroller.Search)
	r.GET("/articles/:article", articlecontroller.Show)
	r.GET("/comments", commentcontroller.Search)
	r.GET("/tags", tagcontroller.Search)

	// 加载验证中间件
	r.Use(middleware.Auth())
	// 用户接口
	r.GET("/users", usercontroller.Search)
	r.GET("/users/:user", usercontroller.Show)
	r.PUT("/users/:user/update", usercontroller.Update)
	// 文章接口
	r.PUT("/articles/:article", articlecontroller.Update)
	r.DELETE("/articles/:article", articlecontroller.Delete)
	r.POST("/articles/:article/like", articlecontroller.Like)
	r.POST("/articles/:article/collect", articlecontroller.Collect)
	r.POST("/articles/:article/comment", articlecontroller.Comment)
	// 评论接口
	r.PUT("/comments/:comment", commentcontroller.Update)
	r.DELETE("/comments/:comment", commentcontroller.Delete)
	r.POST("/comments/:comment/like", commentcontroller.Like)
	r.POST("/comments/:comment/reply", commentcontroller.Reply)
	// 标签接口
	r.POST("/tags/tag/:model", tagcontroller.Tag)
	r.PUT("/tags/:tag", tagcontroller.Update)
	r.DELETE("/tags/:tag", tagcontroller.Delete)

	// 启动服务
	r.Run()
}
