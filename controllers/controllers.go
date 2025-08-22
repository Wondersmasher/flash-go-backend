package controllers

import (
	"errors"

	u "github.com/flash-backend/utils"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	u.AppResponse(c, 200, nil, "Server is healthy, alive and running!")

}

func NoRoute(c *gin.Context) {
	u.AppError(c, 404, errors.New("route not found"), "Route not found")
}
