package routes

import (
	"github.com/MattCSN/project-par/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	golfGroup := r.Group("/golfs")
	{
		golfGroup.GET("/", controllers.GetGolfs)
		golfGroup.POST("/", controllers.CreateGolf)
		golfGroup.POST("/list", controllers.AddGolfs)
		golfGroup.PATCH("/:id", controllers.UpdateGolf)
		golfGroup.DELETE("/:id", controllers.DeleteGolf)
	}
	return r
}
