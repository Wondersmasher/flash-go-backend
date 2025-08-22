package routes

import (
	contr "github.com/flash-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(server *gin.RouterGroup) {
	authGroup := server.Group("/auth")

	authGroup.POST("/sign-up", contr.SignUp)
	authGroup.POST("/sign-in", contr.SignIn)
	authGroup.POST("/sign-out", contr.SignOut)
}
