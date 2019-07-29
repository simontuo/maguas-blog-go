package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"maguas-blog-go/config"
	"maguas-blog-go/log"
	"maguas-blog-go/router"
)

func main() {
	// 日志
	log.Log()

	// 路由实例
	r := gin.Default()
	// 初始化路由
	router.InitRouter(r)
	// 启动服务
	r.Run(fmt.Sprintf(":%v", config.Port))
}
