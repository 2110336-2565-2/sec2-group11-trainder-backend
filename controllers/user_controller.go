package controllers

import (
	"net/http"

	"trainder-api/models"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

type ProfileDetails struct {
	FirstName   string `json:"firstname" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	BirthDate   string `json:"birthdate" binding:"required"`
	CitizenId   string `json:"citizenId" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address     string `json:"address" binding:"required"`
	AvatarUrl   string `json:"avatarUrl"`
}

// CurrentUser godoc
//
//	@Summary		get the current user's username
//	@Description	get the current user's username.  After getting token replied from logging in, put token in ginContext's token field
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	responses.CurrentUserResponse
//	@Failure		400	{object}	responses.CurrentUserResponse
//	@Security		BearerAuth
//	@Router			/protected/user [get]
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

// UpdateProfile godoc
//
//	@Summary		updateProfile of the current user
//	@Description	updateProfile of the current user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			ProfileToUpdate	body		ProfileDetails	true	"put profile input json and pass to  gin.Context"
//	@Success		200				{object}	responses.ProfileResponse
//	@Failure		400				{object}	responses.ProfileResponse
//	@Security		BearerAuth
//	@Router			/protected/update-profile [post]
func UpdateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ProfileDetails
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		err = models.ProfileConditionCheck(input.FirstName, input.LastName, input.BirthDate, input.CitizenId, input.Gender, input.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		_, err = models.UpdateUserProfile(username, input.FirstName, input.LastName, input.BirthDate, input.CitizenId, input.Gender, input.PhoneNumber, input.Address, input.AvatarUrl)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{
				Status:  http.StatusBadRequest,
				Message: `update failed`,
			})
			return
		}

		c.JSON(http.StatusOK,
			responses.ProfileResponse{
				Status:  http.StatusOK,
				Message: username + ` update success!`,
			})
	}
}

// getProfile godoc
//
//	@Summary		getProfile of the current user
//	@Description	getProfile of the current user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.UserProfileResponse
//	@Failure		401	{object}	responses.UserProfileResponse	"Unauthorized, the user is not logged in"
//	@Security		BearerAuth
//	@Router			/protected/profile [get]
func GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserProfileResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		result, err := models.FindProfile(username, "")
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.UserProfileResponse{
				Status:  http.StatusBadRequest,
				Message: `User profile retrieval unsuccessful`,
			})
			return
		}

		c.JSON(http.StatusOK, responses.UserProfileResponse{
			Status:  http.StatusOK,
			Message: `Successfully retrieve user profile`,
			User:    result,
		})
		_ = result
	}
}
