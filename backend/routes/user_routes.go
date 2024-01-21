package routes

import (
	"github.com/arioprima/jobseeker/tree/main/backend/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(authController *controller.AuthController, adminController *controller.AdminController) *gin.Engine {
	service := gin.Default()

	router := service.Group("/api")

	router.POST("/auth/login", authController.Login)
	router.POST("/auth/register", authController.Register)
	router.POST("/auth/verify-email", authController.VerifyEmail)

	router.POST("/admin/save", adminController.Save)
	router.POST("/admin/update", adminController.Update)

	return service
}
