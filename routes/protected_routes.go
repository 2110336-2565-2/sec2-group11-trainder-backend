package routes

import (
	"trainder-api/controllers"
	"trainder-api/middlewares"

	"github.com/gin-gonic/gin"
)

func ProtectedRoute(router *gin.Engine) {
	protected := router.Group("/protected")
	protected.Use(middlewares.JwtAuthMiddleware())

	// Current user information
	protected.GET("/user", controllers.CurrentUser())
	protected.GET("/profile", controllers.GetProfile())

	// Own Trainer Information
	protected.GET("/trainer-profile", controllers.CurrentTrainerUserProfile())

	// Update data
	protected.POST("/update-profile", controllers.UpdateProfile())
	protected.POST("/update-trainer", controllers.UpdateTrainerProfile())

	// Get Others Trainer information
	protected.POST("/trainer", controllers.GetTrainerProfile())
	protected.POST("/filter-trainer", controllers.FilterTrainer())

	// Get review
	protected.POST("/get-reviews", controllers.GetReviews())

	// Add review
	protected.POST("/add-review", controllers.AddTrainerReview())

	protected.POST("/create-booking",)
}
