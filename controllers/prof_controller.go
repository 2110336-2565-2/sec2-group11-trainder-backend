package controllers

import (
	"fmt"
	"net/http"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

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

		c.JSON(http.StatusCreated,
			responses.ProfileResponses{
				Status:  http.StatusCreated,
				Message: Username + ` update success!`,
			})
	}
}
