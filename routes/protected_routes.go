package routes

import (
	"trainder-api/controllers"
	"trainder-api/middlewares"
	"trainder-api/ws"

	"github.com/gin-gonic/gin"
)

func ProtectedRoute(router *gin.Engine, wsHandler *ws.Handler) {
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
	protected.POST("/reviews", controllers.GetReviews())

	// Add review
	protected.POST("/add-review", controllers.AddTrainerReview())

	// Add booking
	protected.POST("/create-booking", controllers.Book())

	// Get bookings
	protected.GET("/bookings", controllers.GetBookings())
	protected.GET("/today-event", controllers.GetTodayEvents())

	// Booking
	protected.POST("/update-booking", controllers.UpdateBooking())
	protected.DELETE("/delete-booking", controllers.DeleteBooking())

	// Reviewable
	protected.POST("/reviewable", controllers.Reviewable())

	// chat
	protected.POST("/create-room", wsHandler.CreateRoom)
	protected.GET("/get-rooms", wsHandler.GetRooms)
	protected.GET("/get-clients/:roomId", wsHandler.GetClients)
	router.GET("/join-room/:roomId", wsHandler.JoinRoom)

	//chat and DB part
	protected.GET("/get-RoomID", controllers.GetRoomID())
	protected.GET("/get-All-Chats", controllers.GetChatsAndLatestMessege())
	protected.GET("/get-past-chat", controllers.GetPastChat())

}
