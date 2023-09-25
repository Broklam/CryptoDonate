package server

import (
	"github.com/Broklam/cryptodonate/backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/searchName", controllers.FindStreamerByName)

	}
	return router
}
