package routes

import (
	"zoop/one/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/urls/register", controllers.RegisterEndPoints)

	router.GET("/proxy", controllers.GetData)
	router.POST("/proxy", controllers.GetData)
	router.PUT("/proxy", controllers.GetData)
	router.PATCH("/proxy", controllers.GetData)
	router.DELETE("/proxy", controllers.GetData)

	return router
}
