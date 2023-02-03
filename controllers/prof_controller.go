package controllers

import (
	"trainder-api/configs"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

type ProfileInput struct {
	Firstname   string `json:"firstname" binding:"required"`
	Lastname    string `json:"lastname" binding:"required"`
	Birthdate   string `jason:"birthdate" binding:"required"`
	CitizenId   string `jason:"citizenid" binding:"required"`
	Gender      string `jason:"gender" binding:"required"`
	PhoneNumber string `jason:"phonenumber" binding:"required"`
	Address     string `json:"addresss" binding:"required"`
	SubAddress  string `json:"subaddresss" binding:"required"`
}

func UpdateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("test")
		var input ProfileInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponses{
				Status:  http.StatusBadRequest,
				Message: err.Error()})
			return
		}
		Username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		//update mongo
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result,err := userCollection.UpdateOne(
			ctx,
			bson.M{"username": Username},
			bson.M{"$set": bson.M{
				"firstname":   input.Firstname,
				"lastname":    input.Lastname,
				"birthdate":   input.Birthdate,
				"citizenid":   input.CitizenId,
				"gender":      input.Gender,
				"phonenumber": input.PhoneNumber,
				"address":     input.Address,
				"subaddress":  input.SubAddress,
			}},
		)
		_ = result
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadRequest,
				Message: `update failed`,
			})
			return
		}

		c.JSON(http.StatusCreated,
			responses.ProfileResponses{
				Status:  http.StatusCreated,
				Message: Username + ` update success!`,
			})
	}
}
