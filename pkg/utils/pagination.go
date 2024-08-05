package utils

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaginationParams(c *gin.Context) (int, int) {
	page, err := strconv.Atoi(c.DefaultQuery("page", os.Getenv("DEFAULT_PAGE")))
	if err != nil {
		page, _ = strconv.Atoi(os.Getenv("DEFAULT_PAGE"))
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", os.Getenv("DEFAULT_PAGE_SIZE")))
	if err != nil {
		pageSize, _ = strconv.Atoi(os.Getenv("DEFAULT_PAGE_SIZE"))
	}
	return page, pageSize
}
