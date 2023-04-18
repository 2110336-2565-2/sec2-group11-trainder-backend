package routes

import (
	"trainder-api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/register", controllers.Register())
	router.POST("/login", controllers.Login())

}
