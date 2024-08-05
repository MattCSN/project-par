package view

import (
	"github.com/MattCSN/project-par/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetGolfWithDetails Get golf details
// @Summary Get golf details
// @Description Get detailed information about a specific golf entity, including courses, holes, and tees
// @Tags Views
// @Param golf_id path string true "Golf ID"
// @Router /v1/golfs/{golf_id}/details [get]
func GetGolfWithDetails(ctx *gin.Context) {
	golfID := ctx.Param("golf_id")
	golfDetails, err := viewService.GetGolfWithDetails(golfID)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, golfDetails)
}
