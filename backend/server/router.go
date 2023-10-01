package server

import (
	"github.com/Broklam/cryptodonate/backend/controllers"
	middlewares "github.com/Broklam/cryptodonate/backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/user/register/streamer/full", controllers.RegisterFullStreamer)
		api.POST("/donations/create", controllers.CreateDonation)

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Tick)
		}

		//secured := api.Group("/secured").Use(middlewares.Auth())
		//{
		//	secured.GET("/ping", controllers.Tick)
		//}
		//}
	}
	return router
}
