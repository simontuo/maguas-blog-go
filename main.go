package main

import (
	"github.com/gin-gonic/gin"
	"maguas-blog/controller"
	"maguas-blog/middleware"
)

func main() {
	r := gin.Default()
	r.GET("/", controller.Index)

	// 加载验证中间件
	r.Use(middleware.Auth())
	r.GET("/test", controller.Test)

	r.Run()
}
