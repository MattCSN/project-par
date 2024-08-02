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
	golfGroup.GET("/:golf_id", golf.GetGolfByID)
	golfGroup.POST("/", golf.CreateGolf)
	golfGroup.PATCH("/:golf_id", golf.UpdateGolf)
	golfGroup.DELETE("/:golf_id", golf.DeleteGolf)
	golfGroup.GET("/:golf_id/courses", course.GetCoursesByGolfID) // Changed :golfId to :golf_id

	// Course router
	courseGroup := v1.Group("/courses")
	courseGroup.GET("/", course.GetCourses)
	courseGroup.GET("/:course_id", course.GetCourseByID) // Changed :id to :course_id
	courseGroup.POST("/", course.CreateCourse)
	courseGroup.PATCH("/:course_id", course.UpdateCourse)  // Changed :id to :course_id
	courseGroup.DELETE("/:course_id", course.DeleteCourse) // Changed :id to :course_id

	// Hole router
	holeGroup := v1.Group("/holes")
	holeGroup.GET("/", hole.GetHoles)
	holeGroup.GET("/:hole_id", hole.GetHoleByID) // Changed :id to :hole_id
	holeGroup.POST("/", hole.CreateHole)
	holeGroup.PATCH("/:hole_id", hole.UpdateHole)  // Changed :id to :hole_id
	holeGroup.DELETE("/:hole_id", hole.DeleteHole) // Changed :id to :hole_id

	// Tee router
	teeGroup := v1.Group("/tees")
	teeGroup.GET("/", tee.GetTees)
	teeGroup.GET("/:tee_id", tee.GetTeeByID) // Changed :id to :tee_id
	teeGroup.POST("/", tee.CreateTee)
	teeGroup.PATCH("/:tee_id", tee.UpdateTee)  // Changed :id to :tee_id
	teeGroup.DELETE("/:tee_id", tee.DeleteTee) // Changed :id to :tee_id

	return r
}
