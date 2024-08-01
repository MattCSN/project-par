package tee

import (
	"github.com/MattCSN/project-par/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTees gets all tees
// @Summary Get all tees
// @Tags Tees
// @Produce json
// @Success 200 {array} tee.Model
// @Failure 500 {object} AppError
// @Router /v1/tees [get]
func GetTees(c *gin.Context) {
	tees, err := teeService.GetAllTees()
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
// @Param tee body tee.Model true "Tee"
// @Success 201 {object} tee.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/tees [post]
func CreateTee(c *gin.Context) {
	var tee Tee
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
// @Param id path string true "Tee ID"
// @Param tee body tee.Model true "Tee"
// @Success 200 {object} tee.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/tees/{id} [patch]
func UpdateTee(c *gin.Context) {
	teeID := c.Param("id")
	var tee Tee
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
// @Param id path string true "Tee ID"
// @Success 204 "No Content"
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/tees/{id} [delete]
func DeleteTee(c *gin.Context) {
	teeID := c.Param("id")
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
// @Param id path string true "Tee ID"
// @Success 200 {object} tee.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/tees/{id} [get]
func GetTeeByID(c *gin.Context) {
	id := c.Param("id")
	tee, err := teeService.GetTeeByID(id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, tee)
}
