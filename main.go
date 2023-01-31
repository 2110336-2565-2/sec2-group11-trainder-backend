package main

import (
	"net/http"
	"trainder-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello software engineering 2",
		})
	})
	public := r.Group("/api")
	public.POST("/register", controllers.Register)

	r.Run(":8080")
}
