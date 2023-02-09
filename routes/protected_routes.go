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
	protected.POST("/update-profile", controllers.UpdateProfile())
	protected.GET("/profile", controllers.GetProfile())
	protected.GET("/trainer", controllers.GetTrainer())
	protected.POST("/filter-trainer", controllers.FilterTrainer())
	protected.POST("/update-trainer", controllers.UpdateTrainer())
}
