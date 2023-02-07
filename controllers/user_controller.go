package controllers

import (
	"net/http"

	"trainder-api/models"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

type ProfileInput struct {
	FirstName   string `json:"firstname" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	BirthDate   string `json:"birthdate" binding:"required"`
	CitizenId   string `json:"citizenId" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address     string `json:"address" binding:"required"`
	SubAddress  string `json:"subAddress" binding:"required"`
}

type GetTrainerInput struct {
	Username string `json:"username" binding:"required"`
}

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
		err = models.ProfileConditionCheck(input.FirstName, input.LastName, input.BirthDate, input.CitizenId, input.Gender, input.PhoneNumber)
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

		result, err := models.FindProfile(username, "")
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadRequest,
				Message: `User profile retrieval unsuccessful`,
			})
			return
		}

		c.JSON(http.StatusOK, responses.GetProfileResponses{
			Status:  http.StatusOK,
			Message: `Successfully retrieve user profile`,
			User:    result,
		})
		_ = result
	}
}

func GetTrainer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input GetTrainerInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.GetTrainerResponses{
				Status:  http.StatusBadRequest,
				Message: "input missing"})
			return
		}
		result, err := models.FindProfile(input.Username, "trainer")
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadRequest,
				Message: `trainer profile retrieval unsuccessful`,
			})
			return
		}
		c.JSON(http.StatusFound, responses.GetProfileResponses{
			Status:  http.StatusFound,
			Message: `Successfully retrieve trainer profile`,
			User:    result,
		})
	}
}
