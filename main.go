package main

import (
	// "fmt"
	"net/http"
	// "trainder-api/configs"
	"trainder-api/routes"

	_ "trainder-api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Trainder API
// @version		0.1
// @description	API for Trainder
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
	//nat edited
	// router.POST("/udprof", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": "xxxxxx Welcome to Trainder API",
	// 	})
	// })
	// var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "usersx")
	// fmt.Println("printDB", *configs.DB)
	// fmt.Println("printuser", userCollection.Name(), userCollection.Database().Name())
	// router.POST("/nao", controllers.UpdateProfile())
	// router.GET("/nao", controllers.GetProfile())

	routes.AuthRoute(router)
	routes.ProtectedRoute(router)
	// router.Run(":8080")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
<<<<<<< HEAD
	// router.Run(":8080")
	router.Run("127.0.0.1:8080")
=======
	router.Run(":8080")
	// router.Run("127.0.0.1:8080")
>>>>>>> 3ea37127235db70002246fa4fb7584dea51eb65e
}
