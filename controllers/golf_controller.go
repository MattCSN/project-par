package controllers

import (
	"github.com/MattCSN/project-par/models"
	"github.com/MattCSN/project-par/repository"
	"github.com/MattCSN/project-par/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var golfService = services.NewGolfService(repository.NewGolfRepository())

func GetGolfs(c *gin.Context) {
	if golfs, err := golfService.GetAllGolfs(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, golfs)
	}
}

func CreateGolf(c *gin.Context) {
	var golf models.Golf
	if err := c.ShouldBindJSON(&golf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := golfService.CreateGolf(&golf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "Golf created successfully"})
	}
}

func AddGolfs(c *gin.Context) {
	var golfs []models.Golf
	if err := c.ShouldBindJSON(&golfs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := golfService.AddGolfs(golfs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "Golfs added successfully"})
	}
}
