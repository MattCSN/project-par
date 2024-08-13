package course

import (
	"github.com/MattCSN/project-par/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GetCourses gets all courses with pagination
// @Summary Get all courses
// @Description Get all courses with pagination
// @Tags Courses
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param pageSize query int false "Page size (default is 10)"
// @Success 200 {array} course.Model
// @Failure 500 {object} AppError
// @Router /v1/courses [get]
func GetCourses(c *gin.Context) {
	page, pageSize := utils.GetPaginationParams(c)

	courses, err := courseService.GetAllCourses(page, pageSize)
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
	var course Model
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

// UpdateCourse updates an existing course by course_id
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
	courseID := c.Param("course_id")
	var course Model
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

// DeleteCourse deletes a course by course_id
// @Summary Delete a course by ID
// @Tags Courses
// @Param id path string true "Course ID"
// @Success 204 "No Content"
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/courses/{id} [delete]
func DeleteCourse(c *gin.Context) {
	courseID := c.Param("course_id")
	if err := courseService.DeleteCourse(courseID); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// GetCourseByID gets a course by course_id
// @Summary Get a course by ID
// @Tags Courses
// @Produce json
// @Param id path string true "Course ID"
// @Success 200 {object} course.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/courses/{id} [get]
func GetCourseByID(c *gin.Context) {
	courseID := c.Param("course_id")
	course, err := courseService.GetCourseByID(courseID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, course)
}

// GetCoursesByGolfID gets all courses for a specific golf with pagination
// @Summary Get all courses for a specific golf
// @Tags Courses
// @Produce json
// @Param golf_id path string true "Golf ID"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} course.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/golfs/{golf_id}/courses [get]
func GetCoursesByGolfID(c *gin.Context) {
	golfID := c.Param("golf_id")
	page, pageSize := utils.GetPaginationParams(c)

	courses, err := courseService.GetCoursesByGolfID(golfID, page, pageSize)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, courses)
}

// GetCoursesByGolfIDs gets all courses for multiple golfs with pagination
// @Summary Get all courses for multiple golfs
// @Tags Courses
// @Produce json
// @Param golf_ids query []string true "Golf IDs"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} course.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/golfs/courses [get]
func GetCoursesByGolfIDs(c *gin.Context) {
	golfIDs := c.Query("golf_ids")
	page, pageSize := utils.GetPaginationParams(c)

	// Split the golfIDs into individual UUIDs
	validGolfIDs := strings.Split(golfIDs, ",")

	courses, err := courseService.GetCoursesByGolfIDs(validGolfIDs, page, pageSize)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, courses)
}

// GetCourseDetails godoc
// @Summary Get detailed information about all courses
// @Description Get detailed information about all courses, including ID, creation date, last update date, course name, golf name, number of holes, pitch and putt status, compact status, postal code, city, and country
// @Tags Courses
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param pageSize query int false "Page size (default is 10)"
// @Success 200 {array} course.DetailsDTO
// @Failure 500 {object} AppError
// @Router /v1/courses/details [get]
func GetCourseDetails(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	courseDetails, err := courseService.GetCourseDetails(page, pageSize)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, courseDetails)
}

// SearchCourseDetails searches for course details by name, city, postal code, or proximity to given coordinates with pagination
// @Summary Search course details by name, city, postal code, or proximity to given coordinates
// @Description Search course details by name, city, postal code, or proximity to given coordinates with pagination
// @Tags Courses
// @Produce json
// @Param searchTerm query string false "Search term"
// @Param latitude query float64 false "Latitude"
// @Param longitude query float64 false "Longitude"
// @Param page query int false "Page number (default is 1)"
// @Param pageSize query int false "Page size (default is 10)"
// @Success 200 {array} course.DetailsDTO
// @Failure 500 {object} AppError
// @Router /v1/courses/details/search [get]
func SearchCourseDetails(c *gin.Context) {
	searchTerm := c.Query("searchTerm")
	longitude, _ := strconv.ParseFloat(c.Query("longitude"), 64)
	latitude, _ := strconv.ParseFloat(c.Query("latitude"), 64)
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	var courseDetails []DetailsDTO

	if searchTerm != "" {
		courseDetails, err = courseService.SearchCourseDetails(searchTerm, page, pageSize)
	} else if longitude != 0 && latitude != 0 {
		courseDetails, err = courseService.SearchCourseDetailsByProximity(longitude, latitude, page, pageSize)
	} else {
		utils.HandleError(c, utils.NewAppError(http.StatusBadRequest, "Invalid search parameters"))
		return
	}

	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, courseDetails)
}
