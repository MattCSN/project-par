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

	// Golf router
	golfGroup := v1.Group("/golfs")
	golfGroup.GET("/search", golf.SearchGolfs) // Register the search route first
	golfGroup.GET("/", golf.GetGolfs)
	golfGroup.GET("/:golf_id", golf.GetGolfByID)
	golfGroup.POST("/", golf.CreateGolf)
	golfGroup.PATCH("/:golf_id", golf.UpdateGolf)
	golfGroup.DELETE("/:golf_id", golf.DeleteGolf)
	golfGroup.GET("/:golf_id/courses", course.GetCoursesByGolfID)

	// Course router
	courseGroup := v1.Group("/courses")
	courseGroup.GET("/", course.GetCourses)
	courseGroup.GET("/:course_id", course.GetCourseByID)
	courseGroup.POST("/", course.CreateCourse)
	courseGroup.PATCH("/:course_id", course.UpdateCourse)
	courseGroup.DELETE("/:course_id", course.DeleteCourse)
	courseGroup.GET("/:course_id/holes", hole.GetHolesByCourseID)

	// Hole router
	holeGroup := v1.Group("/holes")
	holeGroup.GET("/", hole.GetHoles)
	holeGroup.GET("/:hole_id", hole.GetHoleByID)
	holeGroup.POST("/", hole.CreateHole)
	holeGroup.PATCH("/:hole_id", hole.UpdateHole)
	holeGroup.DELETE("/:hole_id", hole.DeleteHole)
	holeGroup.GET("/:hole_id/tees", tee.GetTeesByHoleID)

	// Tee router
	teeGroup := v1.Group("/tees")
	teeGroup.GET("/", tee.GetTees)
	teeGroup.GET("/:tee_id", tee.GetTeeByID)
	teeGroup.POST("/", tee.CreateTee)
	teeGroup.PATCH("/:tee_id", tee.UpdateTee)
	teeGroup.DELETE("/:tee_id", tee.DeleteTee)

	return r
}
