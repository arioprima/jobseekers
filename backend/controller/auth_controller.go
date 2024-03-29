package controller

import (
	"fmt"
	"net/http"

	"github.com/arioprima/jobseeker/tree/main/backend/models"
	"github.com/arioprima/jobseeker/tree/main/backend/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	loginRequest := models.LoginInput{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}

	loginResponse, err := controller.AuthService.Login(ctx, loginRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": fmt.Sprintf("%v", err),
		})
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    loginResponse,
		})
	}
}

func (controller *AuthController) Register(ctx *gin.Context) {
	registerRequest := models.RegisterInput{}
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	registerResponse, err := controller.AuthService.Register(ctx, registerRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("%v", err),
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "OK",
			"data":    registerResponse,
		})
	}
}

func (controller *AuthController) VerifyEmail(ctx *gin.Context) {
	verifyRequest := models.VerifyInput{}
	if err := ctx.ShouldBindJSON(&verifyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verifyResponse, err := controller.AuthService.VerifyEmail(ctx, verifyRequest)
	if err != nil {
		// log.Printf("Error verifying email: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("%v", err), "status": http.StatusInternalServerError})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "OK",
		"data":    verifyResponse,
	})
}
