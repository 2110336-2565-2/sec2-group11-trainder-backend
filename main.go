package main

import (
	// "fmt"
	"net/http"
	// "trainder-api/configs"
	"trainder-api/routes"

	"github.com/gin-gonic/gin"
	// "trainder-api/controllers" //nat edited
	// "go.mongodb.org/mongo-driver/bson"//nat edited
	// "go.mongodb.org/mongo-driver/mongo" //nat edited
)

func main() {
	router := gin.Default()
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
	router.Run("127.0.0.1:8080")
}
