package server

import (
	"github.com/Broklam/cryptodonate/backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/user/register/streamer/full", controllers.RegisterFullStreamer)

		//secured := api.Group("/secured").Use(middlewares.Auth())
		//{
		//	secured.GET("/ping", controllers.Tick)
		//}
		//}
	}
	return router
}
