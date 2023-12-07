package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	//初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return nil
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/bluebell/v1")
	//登录和注册业务
	v1.POST("/login", controller.UserLoginHandler)
	v1.POST("/signup", controller.UserSignUpHandler)
	v1.GET("/refresh/token", controller.RefreshTokenHandler) //刷新token

	//帖子业务
	v1.GET("/post/:id", controller.GetPostDetailHandler)
	v1.GET("/posts", controller.GetPostListHandler)
	v1.GET("/posts/order", controller.GetPostListHandler2)
	v1.GET("/posts/search", controller.PostSearchHandler) //帖子搜索

	//社区业务
	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	v1.GET("/community/order", controller.GetCommunityPostListHandler) //根据社区id
	//Github热榜
	//...........待完成
	//评论列表
	v1.GET("/comment/list", controller.CommentListHandler)
	//中间件
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		//创建帖子
		v1.POST("/post/create", controller.CreatePostHandler)
		//投票功能
		v1.POST("/vote", controller.PostVoteHandler)
		//创建评论
		v1.POST("/comment", controller.CommentHandler)

		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})
	return r
}
