package main

import (
	"net/http"
	"trainder-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Welcome to Trainder API",
		})
	})

	routes.AuthRoute(router)
	routes.ProtectedRoute(router)
	router.Run(":8080")
}
