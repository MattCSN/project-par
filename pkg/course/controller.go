package course

import (
	"github.com/MattCSN/project-par/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCourses gets all courses
// @Summary Get all courses
// @Tags Courses
// @Produce json
// @Success 200 {array} course.Model
// @Failure 500 {object} AppError
// @Router /v1/courses [get]
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
// @Tags Courses
// @Accept json
// @Produce json
// @Param course body course.Model true "Course"
// @Success 201 {object} course.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/courses [post]
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
// @Summary Update a course by ID
// @Tags Courses
// @Accept json
// @Produce json
// @Param id path string true "Course ID"
// @Param course body course.Model true "Course"
// @Success 200 {object} course.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/courses/{id} [patch]
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

// DeleteCourse deletes a course by ID
// @Summary Delete a course by ID
// @Tags Courses
// @Param id path string true "Course ID"
// @Success 204 "No Content"
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/courses/{id} [delete]
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
// @Tags Courses
// @Produce json
// @Param id path string true "Course ID"
// @Success 200 {object} course.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/courses/{id} [get]
func GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	course, err := courseService.GetCourseByID(id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, course)
}