package controllers

import (
	"net/http"

	"trainder-api/responses"
	"trainder-api/utils/tokens"
	"trainder-api/models"

	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadGateway,
				Message: err.Error(),
			})
		}
		c.JSON(http.StatusOK, responses.CurrentUserResponse{
			Status:   http.StatusOK,
			Username: username,
		})
	}

}

type ProfileInput struct {
	FirstName   string `json:"firstname" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	BirthDate   string `json:"birthdate" binding:"required"`
	CitizenId   string `json:"citizenid" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	PhoneNumber string `json:"phonenumber" binding:"required"`
	Address     string `json:"addresss" binding:"required"`
	SubAddress  string `json:"subaddresss" binding:"required"`
}

func UpdateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ProfileInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponses{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		_, err = models.UpdateUserProfile(username, input.FirstName, input.LastName, input.BirthDate, input.CitizenId, input.Gender, input.PhoneNumber, input.Address, input.SubAddress)
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
				Message: username + ` update success!`,
			})
	}
}

func GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		result, err := models.FindProfile(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadRequest,
				Message: `cannot find this user profile`,
			})
			return
		}

		c.JSON(http.StatusOK, responses.GetProfileResponses{
			Status:      http.StatusOK,
			ProfileInfo: result,
		})
	}
}

