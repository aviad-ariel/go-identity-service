package controllers

import (
	"github.com/gin-gonic/gin"
	"go-identity-service/helpers"
	"go-identity-service/models"
	"go-identity-service/services"
	"net/http"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SingUp(context *gin.Context) {
	var userService = services.UserService{}
	var user models.User

	jsonError := context.ShouldBindJSON(&user)
	if jsonError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": jsonError.Error()})
		context.Abort()
		return
	}

	validationMessage, validationStatus := ValidateUser(&user)
	if validationStatus != http.StatusOK {
		context.JSON(validationStatus, gin.H{"error": validationMessage})
		context.Abort()
		return
	}

	err := userService.Create(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": jsonError.Error()})
		context.Abort()
		return
	}
}

func SingIn(context *gin.Context) {
	var request TokenRequest
	var userService = services.UserService{}

	jsonError := context.ShouldBindJSON(&request)
	if jsonError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": jsonError.Error()})
		context.Abort()
		return
	}

	// check if email exists and password is correct
	user, err := userService.GetOne("email", request.Email)
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"error": "Invalid credentials"})
		context.Abort()
		return
	}
	credentialError := helpers.CheckPassword(user.Password, request.Password)
	if credentialError != nil || err != nil {
		context.JSON(http.StatusForbidden, gin.H{"error": "Invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := helpers.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"token": tokenString})
}

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
