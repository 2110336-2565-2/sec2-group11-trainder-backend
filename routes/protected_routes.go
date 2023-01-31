package routes

import (
	"trainder-api/controllers"
	"trainder-api/middlewares"

	"github.com/gin-gonic/gin"
)

func ProtectedRoute(router *gin.Engine) {
	protected := router.Group("/protected")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser())
}
