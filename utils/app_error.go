package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppError struct {
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(statusCode int, message string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func HandleAppError(c *gin.Context, err *AppError) {
	c.JSON(err.StatusCode, gin.H{"error": err.Error()})
}

func NotFoundError(resource string) *AppError {
	return NewAppError(http.StatusNotFound, fmt.Sprintf("%s not found", resource))
}

func ConflictError(resource string) *AppError {
	return NewAppError(http.StatusConflict, fmt.Sprintf("%s must be unique", resource))
}
