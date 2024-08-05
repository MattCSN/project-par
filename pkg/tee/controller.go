package tee

import (
	"github.com/MattCSN/project-par/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTees gets all tees with pagination
// @Summary Get all tees
// @Description Get all tees with pagination
// @Tags Tees
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param pageSize query int false "Page size (default is 10)"
// @Success 200 {array} tee.Model
// @Failure 500 {object} AppError
// @Router /v1/tees [get]
func GetTees(c *gin.Context) {
	page, pageSize := utils.GetPaginationParams(c)

	tees, err := teeService.GetAllTees(page, pageSize)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, tees)
}

// CreateTee creates a new tee
// @Summary Create a new tee
// @Tags Tees
// @Accept json
// @Produce json
// @Param tee body tee.Model true "Model"
// @Success 201 {object} tee.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/tees [post]
func CreateTee(c *gin.Context) {
	var tee Model
	if err := c.ShouldBindJSON(&tee); err != nil {
		utils.HandleError(c, utils.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := teeService.CreateTee(&tee); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, tee)
}

// UpdateTee updates an existing tee by ID
// @Summary Update an existing tee
// @Tags Tees
// @Accept json
// @Produce json
// @Param id path string true "Model ID"
// @Param tee body tee.Model true "Model"
// @Success 200 {object} tee.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/tees/{id} [patch]
func UpdateTee(c *gin.Context) {
	teeID := c.Param("tee_id")
	var tee Model
	if err := c.ShouldBindJSON(&tee); err != nil {
		utils.HandleError(c, utils.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}
	tee.ID = teeID
	updatedTee, err := teeService.UpdateTee(&tee)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, updatedTee)
}

// DeleteTee deletes an existing tee by ID
// @Summary Delete an existing tee
// @Tags Tees
// @Param id path string true "Model ID"
// @Success 204 "No Content"
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/tees/{id} [delete]
func DeleteTee(c *gin.Context) {
	teeID := c.Param("tee_id")
	if err := teeService.DeleteTee(teeID); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// GetTeeByID gets a tee by ID
// @Summary Get a tee by ID
// @Tags Tees
// @Produce json
// @Param id path string true "Model ID"
// @Success 200 {object} tee.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/tees/{id} [get]
func GetTeeByID(c *gin.Context) {
	id := c.Param("tee_id")
	tee, err := teeService.GetTeeByID(id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, tee)
}

// GetTeesByHoleID gets all tees for a specific hole with pagination
// @Summary Get all tees for a specific hole
// @Tags Tees
// @Produce json
// @Param hole_id path string true "Hole ID"
// @Param page query int false "Page number (default is 1)"
// @Param pageSize query int false "Page size (default is 10)"
// @Success 200 {array} tee.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/holes/{hole_id}/tees [get]
func GetTeesByHoleID(c *gin.Context) {
	holeID := c.Param("hole_id")
	page, pageSize := utils.GetPaginationParams(c)

	tees, err := teeService.GetTeesByHoleID(holeID, page, pageSize)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, tees)
}
