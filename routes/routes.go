package routes

import (
	"github.com/MattCSN/project-par/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Group("/golfs").GET("/", controllers.GetGolfs)
	r.Group("/golfs").POST("/", controllers.CreateGolf)
	r.Group("/golfs").POST("/list", controllers.AddGolfs)
	return r
}
