package router

import (
	"gin_forum/controller"
	"gin_forum/logger"
	"gin_forum/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) 
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")

	v1.POST("/signup", controller.SignUpHandler)

	v1.POST("/login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())
	v1.GET("/posts", controller.GetPostListHandler)
	v1.POST("/createpost", controller.CreatePostHandler)
	v1.GET("/topic/:postid", controller.GetPostDetailHandler)
	v1.POST("/vote",controller.PostVoteController)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
