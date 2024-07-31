package course

import (
	"github.com/MattCSN/project-par/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var courseService = NewCourseService(NewCourseRepository())

// GetCourses gets all courses
// @Summary Get all courses
// @Description Get all courses
// @Tags course
// @Produce json
// @Success 200 {array} Course
// @Failure 500 {object} utils.AppError
// @Router /courses [get]
func GetCourses(c *gin.Context) {
	courses, err := courseService.GetAllCourses()
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, courses)
}

// CreateCourse creates a new course
// @Summary Create a new course
// @Description Create a new course
// @Tags course
// @Accept json
// @Produce json
// @Param course body Course true "Course"
// @Success 201 {object} Course
// @Failure 400 {object} utils.AppError
// @Failure 500 {object} utils.AppError
// @Router /courses [post]
func CreateCourse(c *gin.Context) {
	var course Course
	if err := c.ShouldBindJSON(&course); err != nil {
		utils.HandleError(c, utils.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := courseService.CreateCourse(&course); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, course)
}

// UpdateCourse updates an existing course by ID
// @Summary Update an existing course
// @Description Update an existing course by ID
// @Tags course
// @Accept json
// @Produce json
// @Param id path string true "Course ID"
// @Param course body Course true "Course"
// @Success 200 {object} Course
// @Failure 400 {object} utils.AppError
// @Failure 500 {object} utils.AppError
// @Router /courses/{id} [patch]
func UpdateCourse(c *gin.Context) {
	courseID := c.Param("id")
	var course Course
	if err := c.ShouldBindJSON(&course); err != nil {
		utils.HandleError(c, utils.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}
	course.ID = courseID
	updatedCourse, err := courseService.UpdateCourse(&course)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, updatedCourse)
}

// DeleteCourse deletes an existing course by ID
// @Summary Delete an existing course
// @Description Delete an existing course by ID
// @Tags course
// @Param id path string true "Course ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.AppError
// @Failure 500 {object} utils.AppError
// @Router /courses/{id} [delete]
func DeleteCourse(c *gin.Context) {
	courseID := c.Param("id")
	if err := courseService.DeleteCourse(courseID); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// GetCourseByID gets a course by ID
// @Summary Get a course by ID
// @Description Get a course by ID
// @Tags course
// @Produce json
// @Param id path string true "Course ID"
// @Success 200 {object} Course
// @Failure 400 {object} utils.AppError
// @Failure 500 {object} utils.AppError
// @Router /courses/{id} [get]
func GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	course, err := courseService.GetCourseByID(id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, course)
}
