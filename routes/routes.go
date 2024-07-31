package routes

import (
	"github.com/MattCSN/project-par/course"
	"github.com/MattCSN/project-par/golf"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the API routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		// Golf routes
		golfGroup := v1.Group("/golfs")
		{
			golfGroup.GET("/", golf.GetGolfs)
			golfGroup.GET("/:id", golf.GetGolfByID)
			golfGroup.POST("/", golf.CreateGolf)
			golfGroup.PATCH("/:id", golf.UpdateGolf)
			golfGroup.DELETE("/:id", golf.DeleteGolf)
		}

		// Course routes
		courseGroup := v1.Group("/courses")
		{
			courseGroup.GET("/", course.GetCourses)
			courseGroup.GET("/:id", course.GetCourseByID)
			courseGroup.POST("/", course.CreateCourse)
			courseGroup.PATCH("/:id", course.UpdateCourse)
			courseGroup.DELETE("/:id", course.DeleteCourse)
		}
	}

	return r
}
