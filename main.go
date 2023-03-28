package main

import (
	"net/http"
	"runtime"
	"trainder-api/routes"
	"trainder-api/utils/inits"
	"trainder-api/ws"

	_ "trainder-api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title						Trainder API
// @version					0.1
// @description				API for Trainder
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	inits.InitializeDatabase()
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Origin", "Authorization")
	router.Use(cors.New(config))
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Welcome to Trainder API",
		})
	})

	routes.AuthRoute(router)
	routes.ProtectedRoute(router,wsHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	os := runtime.GOOS
	if os == "windows" {
		router.Run("127.0.0.1:8080")
	} else {
		router.Run(":8080")
	}
}
