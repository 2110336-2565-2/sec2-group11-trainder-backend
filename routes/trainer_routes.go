
package routes

import (
	"trainder-api/controllers"

	"github.com/gin-gonic/gin"
)

func TrainerRoute(router *gin.Engine) {
	router.GET("/trainer", controllers.GetTrainer())
}
