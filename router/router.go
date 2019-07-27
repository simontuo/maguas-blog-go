package router

import (
	"github.com/gin-gonic/gin"
	"maguas-blog-go/controller"
	"maguas-blog-go/middleware/jwt"
)

func InitRouter(r *gin.Engine)  {
	v1 := r.Group("/v1")
	{
		// 获取token接口
		v1.GET("/token", controller.GenerateToken)
		v1.GET("/token/refresh", controller.RefreshToken)
		// 非验证接口
		v1.GET("/", controller.Index)
		v1.GET("/articles", controller.ArticleSearch)
		v1.GET("/articles/:article", controller.ArticleShow)
		v1.GET("/comments", controller.CommentSearch)
		v1.GET("/tags", controller.TagSearch)

		// 加载jwt中间件
		v1.Use(jwt.JWTAuth())

		// 用户接口
		v1.GET("/users", controller.UserSearch)
		v1.POST("/users", controller.UserCreate)
		v1.GET("/users/:user", controller.UserShow)
		v1.PUT("/users/:user/update", controller.UserUpdate)
		v1.DELETE("/users/:user", controller.UserDelete)

		// 文章接口
		v1.PUT("/articles/:article", controller.ArticleUpdate)
		v1.DELETE("/articles/:article", controller.ArticleDelete)
		v1.POST("/articles/:article/like", controller.ArticleLike)
		v1.POST("/articles/:article/collect", controller.ArticleCollect)
		v1.POST("/articles/:article/comment", controller.ArticleComment)

		// 评论接口
		v1.PUT("/comments/:comment", controller.CommentUpdate)
		v1.DELETE("/comments/:comment", controller.CommentDelete)
		v1.POST("/comments/:comment/like", controller.CommentLike)
		v1.POST("/comments/:comment/reply", controller.CommentReply)
		// 标签接口
		v1.POST("/tags/tag/:model", controller.TagTag)
		v1.PUT("/tags/:tag", controller.TagUpdate)
		v1.DELETE("/tags/:tag", controller.TagDelete)
	}
}

