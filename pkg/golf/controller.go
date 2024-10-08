package golf

import (
	"github.com/MattCSN/project-par/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetGolfs gets all golfs with pagination
// @Summary Get all golfs
// @Description Get all golfs with pagination
// @Tags Golfs
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param pageSize query int false "Page size (default is 10)"
// @Success 200 {array} golf.Model
// @Failure 500 {object} AppError
// @Router /v1/golfs [get]
func GetGolfs(ctx *gin.Context) {
	page, pageSize := utils.GetPaginationParams(ctx)

	golfs, err := golfService.GetAllGolfs(page, pageSize)
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
	var golf Model
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
	golfID := ctx.Param("golf_id")
	var golf Model
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
	golfID := ctx.Param("golf_id")
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
	golfID := ctx.Param("golf_id")
	golf, err := golfService.GetGolfByID(golfID)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, golf)
}

// SearchGolfs searches for golfs by name, city, postal code, or proximity to given coordinates with pagination
// @Summary Search golfs by name, city, postal code, or proximity to given coordinates
// @Description Search golfs by name, city, postal code, or proximity to given coordinates with pagination
// @Tags Golfs
// @Produce json
// @Param searchTerm query string false "Search term"
// @Param latitude query float64 false "Latitude"
// @Param longitude query float64 false "Longitude"
// @Param page query int false "Page number (default is 1)"
// @Param pageSize query int false "Page size (default is 10)"
// @Success 200 {array} golf.Model
// @Failure 500 {object} AppError
// @Router /v1/golfs/search [get]
func SearchGolfs(ctx *gin.Context) {
	searchTerm := ctx.Query("searchTerm")
	longitude, _ := strconv.ParseFloat(ctx.Query("longitude"), 64)
	latitude, _ := strconv.ParseFloat(ctx.Query("latitude"), 64)
	page, pageSize := utils.GetPaginationParams(ctx)

	var golfs []Model
	var err error

	if searchTerm != "" {
		golfs, err = golfService.SearchGolfs(searchTerm)
	} else if longitude != 0 && latitude != 0 {
		golfs, err = golfService.SearchGolfsByProximity(longitude, latitude, page, pageSize)
	} else {
		utils.HandleError(ctx, utils.NewAppError(http.StatusBadRequest, "Invalid search parameters"))
		return
	}

	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, golfs)
}
