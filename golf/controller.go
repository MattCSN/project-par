package golf

import (
	"errors"
	"github.com/MattCSN/project-par/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var golfService = NewGolfService(NewGolfRepository())

func handleError(c *gin.Context, err error) {
	var appErr *utils.AppError
	if errors.As(err, &appErr) {
		utils.HandleAppError(c, appErr)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func GetGolfs(c *gin.Context) {
	golfs, err := golfService.GetAllGolfs()
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, golfs)
}

func CreateGolf(c *gin.Context) {
	var golf Golf
	if err := c.ShouldBindJSON(&golf); err != nil {
		handleError(c, utils.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := golfService.CreateGolf(&golf); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Golf created successfully"})
}

func AddGolfs(c *gin.Context) {
	var golfs []Golf
	if err := c.ShouldBindJSON(&golfs); err != nil {
		handleError(c, utils.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := golfService.AddGolfs(golfs); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Golfs added successfully"})
}

func UpdateGolf(c *gin.Context) {
	golfID := c.Param("id")
	var golf Golf
	if err := c.ShouldBindJSON(&golf); err != nil {
		handleError(c, utils.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}
	golf.ID = golfID
	if err := golfService.UpdateGolf(&golf); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Golf updated successfully"})
}

func DeleteGolf(c *gin.Context) {
	golfID := c.Param("id")
	if err := golfService.DeleteGolf(golfID); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Golf deleted successfully"})
}
