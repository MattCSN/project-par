package routes

import (
	"github.com/MattCSN/project-par/golf"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	golfGroup := r.Group("/golfs")
	{
		golfGroup.GET("/", golf.GetGolfs)
		golfGroup.POST("/", golf.CreateGolf)
		golfGroup.POST("/list", golf.AddGolfs)
		golfGroup.PATCH("/:id", golf.UpdateGolf)
		golfGroup.DELETE("/:id", golf.DeleteGolf)
	}
	return r
}
