package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login", gin.H{
		"title": "Login",
	})
}

func (controller *AuthController) SubmitLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if username != "admin" || password != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	session.Set("user", username)
	ctx.Redirect(http.StatusOK, "/admin")
}
