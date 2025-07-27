package routes

import (
	"github.com/gin-gonic/gin"
	"auth-api/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}