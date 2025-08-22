package routes

import (
	contr "github.com/flash-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes(server *gin.Engine) {
	server.NoRoute(contr.NoRoute)

	s := server.Group("/api/v1")

	s.GET("/health-check", contr.HealthCheck)
}
