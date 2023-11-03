package main

import (
	"time"
	"web-app/cofig/constant"
	"web-app/config/constant"
	"web-app/web/auth"
	like "web-app/web/favorite"
	"web-app/web/feed"
	"web-app/web/middleware"
	"web-app/web/user"
	"web-app/web/video"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	g.Use(gzip.Gzip(gzip.DefaultCompression))
	g.Use(middleware.LimitMiddleWare(time.Second, constant.Cap, constant.Quantum))

	g.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	rootPath := g.Group("/app")

	userRouter := rootPath.Group("/user")
	{
		userRouter.GET("/login", middleware.AuthMiddleWare(), auth.LoginHandler)
		userRouter.POST("/register", middleware.AuthMiddleWare(), auth.RegisterHandler)
		userRouter.GET("/info", middleware.AuthMiddleWare(), user.UserInfoHandler)
	}

	feedRouter := rootPath.Group("/feed")
	{
		feedRouter.GET("/feed", feed.FeedHandler)
	}

	favoriteRouter := rootPath.Group("/like")
	{
		favoriteRouter.POST("/do", like.FavoriteHandler)
		favoriteRouter.GET("/list", like.FavoriteListHandler)

	}

	videoRouter := rootPath.Group("/video")
	{
		videoRouter.GET("/list", video.VideoListHandler)
		videoRouter.GET("/tag", video.VideoTagHandler)
	}

	// rootPath
	if err := g.Run(":8080"); err != nil {
		panic(err)
	}

}
