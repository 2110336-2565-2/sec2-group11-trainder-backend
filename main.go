package main

import (
	"net/http"
	"trainder-api/routes"

	_ "trainder-api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Trainder API
//	@version		0.1
//	@description	API for Trainder
func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Origin", "Authorization")
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Welcome to Trainder API",
		})
	})

	routes.AuthRoute(router)
	routes.ProtectedRoute(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
