package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AppError represents an error that can be returned by the application
// @Description Error that can be returned by the application
type AppError struct {
	StatusCode int
	Message    string
} // @Name AppError

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
	return NewAppError(http.StatusNotFound, fmt.Sprintf("NOT FOUND - %s", resource))
}

func ConflictError(resource string) *AppError {
	return NewAppError(http.StatusConflict, fmt.Sprintf("MUST BE UNIQUE - %s", resource))
}

func HandleError(c *gin.Context, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		HandleAppError(c, appErr)
	} else {
		HandleAppError(c, NewAppError(http.StatusInternalServerError, err.Error()))
	}
}
