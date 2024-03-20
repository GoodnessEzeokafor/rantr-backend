package routes

import (
	controllers "rantr/controller"
	"rantr/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.LoginUser)
	router.GET("/auth/user", middleware.AuthMiddleware(), controllers.AuthUser)
	router.POST("/auth/recover-password", controllers.RecoverPassword)
	router.POST("/auth/reset-password", controllers.ResetPassword)
	// Register other routes similarly
	// store user routes
}
