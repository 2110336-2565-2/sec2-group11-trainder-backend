package main

import (
	"net/http"
	"trainder-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Welcome to Trainder API",
		})
	})

	routes.AuthRoute(router)
	router.Run(":8080")
}
