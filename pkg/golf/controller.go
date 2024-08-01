package golf

import (
	"github.com/MattCSN/project-par/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetGolfs gets all golfs
// @Summary Get all golfs
// @Description Get all golfs
// @Tags Golfs
// @Produce json
// @Success 200 {array} golf.Model
// @Failure 500 {object} AppError
// @Router /v1/golfs [get]
func GetGolfs(ctx *gin.Context) {
	golfs, err := golfService.GetAllGolfs()
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, golfs)
}

// CreateGolf creates a new golf
// @Summary Create a new golf
// @Description Create a new golf
// @Tags Golfs
// @Accept json
// @Produce json
// @Param golf body golf.Model true "Golf"
// @Success 201 {object} golf.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/golfs [post]
func CreateGolf(ctx *gin.Context) {
	var golf Golf
	if err := ctx.ShouldBindJSON(&golf); err != nil {
		utils.HandleError(ctx, utils.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := golfService.CreateGolf(&golf); err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, golf)
}

// UpdateGolf updates an existing golf by ID
// @Summary Update an existing golf
// @Description Update an existing golf by ID
// @Tags Golfs
// @Accept json
// @Produce json
// @Param id path string true "Golf ID"
// @Param golf body golf.Model true "Golf"
// @Success 200 {object} golf.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/golfs/{id} [patch]
func UpdateGolf(ctx *gin.Context) {
	golfID := ctx.Param("id")
	var golf Golf
	if err := ctx.ShouldBindJSON(&golf); err != nil {
		utils.HandleError(ctx, utils.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}
	golf.ID = golfID
	updatedGolf, err := golfService.UpdateGolf(&golf)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, updatedGolf)
}

// DeleteGolf deletes an existing golf by ID
// @Summary Delete an existing golf
// @Description Delete an existing golf by ID
// @Tags Golfs
// @Param id path string true "Golf ID"
// @Success 204 "No Content"
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/golfs/{id} [delete]
func DeleteGolf(ctx *gin.Context) {
	golfID := ctx.Param("id")
	if err := golfService.DeleteGolf(golfID); err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// GetGolfByID gets a golf by ID
// @Summary Get a golf by ID
// @Description Get a golf by ID
// @Tags Golfs
// @Produce json
// @Param id path string true "Golf ID"
// @Success 200 {object} golf.Model
// @Failure 400 {object} AppError
// @Failure 500 {object} AppError
// @Router /v1/golfs/{id} [get]
func GetGolfByID(ctx *gin.Context) {
	golfID := ctx.Param("id")
	golf, err := golfService.GetGolfByID(golfID)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, golf)
}
