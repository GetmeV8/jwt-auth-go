package routes

import (
	"go-jwt-auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.GET("/profile", controllers.Profile)
	}
}
