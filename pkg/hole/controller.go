package hole

import (
	"github.com/MattCSN/project-par/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHoles gets all holes
// @Summary Get all holes
// @Tags Holes
// @Produce json
// @Success 200 {array} hole.Model
// @Failure 500 {object} AppError
// @Router /v1/holes [get]
func GetHoles(c *gin.Context) {
	holes, err := holeService.GetAllHoles()
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, holes)
}

// CreateHole creates a new hole
// @Summary Create a new hole
// @Tags Holes
// @Accept json
// @Produce json
// @Param hole body hole.Model true "Hole"
// @Success 201 {object} hole.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/holes [post]
func CreateHole(c *gin.Context) {
	var hole Hole
	if err := c.ShouldBindJSON(&hole); err != nil {
		utils.HandleError(c, utils.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := holeService.CreateHole(&hole); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, hole)
}

// UpdateHole updates an existing hole by ID
// @Summary Update an existing hole
// @Tags Holes
// @Accept json
// @Produce json
// @Param id path string true "Hole ID"
// @Param hole body hole.Model true "Hole"
// @Success 200 {object} hole.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/holes/{id} [patch]
func UpdateHole(c *gin.Context) {
	holeID := c.Param("id")
	var hole Hole
	if err := c.ShouldBindJSON(&hole); err != nil {
		utils.HandleError(c, utils.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}
	hole.ID = holeID
	updatedHole, err := holeService.UpdateHole(&hole)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, updatedHole)
}

// DeleteHole deletes an existing hole by ID
// @Summary Delete an existing hole
// @Tags Holes
// @Param id path string true "Hole ID"
// @Success 204 "No Content"
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/holes/{id} [delete]
func DeleteHole(c *gin.Context) {
	holeID := c.Param("id")
	if err := holeService.DeleteHole(holeID); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// GetHoleByID gets a hole by ID
// @Summary Get a hole by ID
// @Tags Holes
// @Produce json
// @Param id path string true "Hole ID"
// @Success 200 {object} hole.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/holes/{id} [get]
func GetHoleByID(c *gin.Context) {
	id := c.Param("id")
	hole, err := holeService.GetHoleByID(id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, hole)
}