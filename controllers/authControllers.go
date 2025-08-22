package controllers

import (
	"net/http"

	u "github.com/flash-backend/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	u.AppResponse(c, http.StatusCreated, nil, "Sign-up successful")
}

func SignIn(c *gin.Context) {
	u.AppResponse(c, http.StatusOK, nil, "Sign-in successful")
}

func SignOut(c *gin.Context) {
	u.AppResponse(c, http.StatusOK, nil, "Sign-out successful")
}
