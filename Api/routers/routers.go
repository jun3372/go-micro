package routers

import (
	"github.com/gin-gonic/gin"

	"micro.demo/Api/handler"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()

	// 跨域
	router.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")                            // 跨域
		ctx.Header("Access-Control-Allow-Headers", "Token,Content-Type")          // 必须的请求头
		ctx.Header("Access-Control-Allow-Methods", "OPTIONS,PUT,POST,GET,DELETE") // 接收的请求方法
	})

	// OPTIONS
	router.NoRoute(func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			ctx.JSON(200, nil)
		}
	})
}

func Init() *gin.Engine {
	group := router.Group("/v1")
	{
		r := group.Group("/users")
		{
			userHandler := new(handler.UserHandler)
			r.POST("login", userHandler.Login)        // 登录
			r.POST("register", userHandler.Register)  // 注册
			r.PUT("update", userHandler.Update)       // 更新
			r.GET("find", userHandler.Find)           // 查询
			r.GET("findByIds", userHandler.FindByIds) // 查询用户列表
		}
	}
	return router
}
