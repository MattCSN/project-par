package router

import (
	"github.com/MattCSN/project-par/pkg/course"
	"github.com/MattCSN/project-par/pkg/golf"
	"github.com/MattCSN/project-par/pkg/hole"
	"github.com/MattCSN/project-par/pkg/tee"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")

	// Model router
	golfGroup := v1.Group("/golfs")
	golfGroup.GET("/", golf.GetGolfs)
	golfGroup.GET("/:id", golf.GetGolfByID)
	golfGroup.POST("/", golf.CreateGolf)
	golfGroup.PATCH("/:id", golf.UpdateGolf)
	golfGroup.DELETE("/:id", golf.DeleteGolf)

	// Course router
	courseGroup := v1.Group("/courses")
	courseGroup.GET("/", course.GetCourses)
	courseGroup.GET("/:id", course.GetCourseByID)
	courseGroup.POST("/", course.CreateCourse)
	courseGroup.PATCH("/:id", course.UpdateCourse)
	courseGroup.DELETE("/:id", course.DeleteCourse)

	// Hole router
	holeGroup := v1.Group("/holes")
	holeGroup.GET("/", hole.GetHoles)
	holeGroup.GET("/:id", hole.GetHoleByID)
	holeGroup.POST("/", hole.CreateHole)
	holeGroup.PATCH("/:id", hole.UpdateHole)
	holeGroup.DELETE("/:id", hole.DeleteHole)

	// Tee router
	teeGroup := v1.Group("/tees")
	teeGroup.GET("/", tee.GetTees)
	teeGroup.GET("/:id", tee.GetTeeByID)
	teeGroup.POST("/", tee.CreateTee)
	teeGroup.PATCH("/:id", tee.UpdateTee)
	teeGroup.DELETE("/:id", tee.DeleteTee)

	return r
}
