package controller

import (
	"fmt"
	"github.com/arioprima/jobseeker/tree/main/backend/models"
	"github.com/arioprima/jobseeker/tree/main/backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminController struct {
	AdminService service.AdminService
}

func NewAdminController(adminService service.AdminService) *AdminController {
	return &AdminController{AdminService: adminService}
}

func (controller *AdminController) Save(ctx *gin.Context) {
	registerRequest := models.AdminInput{}
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	saveResponse, err := controller.AdminService.Save(ctx, registerRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("%v", err),
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    saveResponse,
		})
	}
}

func (controller *AdminController) Update(ctx *gin.Context) {
	registerRequest := models.AdminInput{}
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updateResponse, err := controller.AdminService.Update(ctx, registerRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("%v", err),
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    updateResponse,
		})
	}
}

func (controller *AdminController) Delete(ctx *gin.Context) {
	deleteRequest := models.User{}
	err := ctx.ShouldBindJSON(&deleteRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	deleteResponse, err := controller.AdminService.Delete(ctx, deleteRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("%v", err),
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    deleteResponse,
		})
	}
}

func (controller *AdminController) FindByID(ctx *gin.Context) {
	findByIDRequest := models.User{}
	err := ctx.ShouldBindJSON(&findByIDRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	findByIDResponse, err := controller.AdminService.FindByID(ctx, findByIDRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("%v", err),
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    findByIDResponse,
		})
	}
}
