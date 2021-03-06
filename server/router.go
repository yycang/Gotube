package server

import (
	"Gotube/api"
	"Gotube/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 视频创建
		v1.POST("video", api.CreateVideo)

		// 视频查看
		v1.GET("video/:id", api.VideoDetail)

		// 视频列表
		v1.GET("videos", api.VideoList)

		// 视频更新
		v1.PUT("video/:id", api.VideoUpdate)

		// 视频删除
		v1.DELETE("video/:id", api.VideoDelete)

		// 需要登录保护的
		v1.Use(middleware.AuthRequired())
		{
			// User Routing
			v1.GET("user/me", api.UserMe)
			v1.DELETE("user/logout", api.UserLogout)
		}

	}
	return r
}
