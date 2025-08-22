package utils

import (
	// "net/http"

	"github.com/gin-gonic/gin"
)

func AppResponse(c *gin.Context, statusCode int, data any, message string) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func AppError(c *gin.Context, statusCode int, err error, message string) {
	c.JSON(statusCode, gin.H{
		"message": "Error",
		"error":   err.Error(),
	})
}
