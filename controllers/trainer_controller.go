package controllers

import (
	"fmt"
	"net/http"

	"trainder-api/models"
	"trainder-api/responses"

	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

type FilterTrainerInput struct {
	Specialty []string `json:"specialty"`
	Limit     int      `json:"limit" binding:"required"`
	FeeMin    float64  `json:"feeMin"`
	FeeMax    float64  `json:"feeMax"`
	// Rating     float64 `json:"Rating" binding:"required"`
}
type UpdateTrainerInput struct {
	Specialty      []string `json:"specialty"`
	Rating         float64  `json:"rating"`
	Fee            float64  `json:"fee"`
	TraineeCount   int32    `json:"traineeCount"`
	CertificateUrl string   `json:"certificateUrl"`
}

// CurrentTrainerUserProfile retrieves the trainer profile of the current user for the user that is a trainer
//
//	@Summary		Retrieve trainer profile of current user
//	@Description	Retrieves the trainer profile information of the current user.
//	@Tags			Trainer
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.TrainerProfileResponse	"Successfully retrieved the trainer profile"
//	@Failure		400	{object}	responses.TrainerProfileResponse	"Failed to retrieve the trainer profile"
//	@Security		BearerAuth
//	@Router			/protected/trainer-profile [get]
func CurrentTrainerUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainerProfileResponse{
				Status:  http.StatusBadRequest,
				Message: "Failed to retrieve the trainer profile",
			})
			return
		}
		userProfile, trainerProfile, err := models.FindTrainerProfile(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainerProfileResponse{
				Status:  http.StatusBadRequest,
				Message: "Failed to retrieve the trainer profile",
			})
			return
		}
		c.JSON(http.StatusOK, responses.TrainerProfileResponse{
			Status:      http.StatusOK,
			Message:     `Successfully retrieve trainer profile`,
			User:        userProfile,
			TrainerInfo: trainerProfile,
		})
	}
}

type GetTrainerInput struct {
	Username string `json:"username" binding:"required"`
}

// GetTrainerProfile retrieves the trainer profile of any trainer
//
//	@Summary		Retrieve trainer profile
//	@Description	Retrieves the trainer profile information.
//	@Tags			Trainer
//	@Accept			json
//	@Produce		json
//	@Param			input	body		GetTrainerInput						true	"Put username input for retrieving the trainer profile"
//	@Success		200		{object}	responses.TrainerProfileResponse	"Successfully retrieved the trainer profile"
//	@Failure		400		{object}	responses.TrainerProfileResponse	"Failed to retrieve the trainer profile"
//	@Security		BearerAuth
//	@Router			/protected/trainer [post]
func GetTrainerProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input GetTrainerInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainerProfileResponse{
				Status:  http.StatusBadRequest,
				Message: "input missing"})
			return
		}
		userProfile, trainerProfile, err := models.FindTrainerProfile(input.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainerProfileResponse{
				Status:  http.StatusBadRequest,
				Message: "Failed to retrieve the trainer profile",
			})
			return
		}
		c.JSON(http.StatusOK, responses.TrainerProfileResponse{
			Status:      http.StatusOK,
			Message:     `Successfully retrieve trainer profile`,
			User:        userProfile,
			TrainerInfo: trainerProfile,
		})
	}
}

// Update the trainer's profile information.
//
//	@Summary	Update the trainer's profile information.
//	@Tags		Trainer
//	@Accept		json
//	@Produce	json
//	@Param		profile	body		UpdateTrainerInput			true	"Trainer's information to update"
//	@Success	200		{object}	responses.ProfileResponse	"Successfully update the trainer's profile"
//	@Failure	400		{object}	responses.ProfileResponse	"Bad Request, either invalid input or user is not a trainer"
//	@Failure	401		{object}	responses.ProfileResponse	"Unauthorized, the user is not logged in"
//	@Security	BearerAuth
//	@Router		/protected/update-trainer [post]
func UpdateTrainerProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input UpdateTrainerInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil || !models.IsTrainer(username) {
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		_, err = models.UpdateTrainerProfile(
			username, input.Specialty, input.Rating,
			input.Fee, input.TraineeCount, input.CertificateUrl)
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

// FilterTrainer godoc
//
//	@Summary		FilterTrainer base on filter input
//	@Description	FilterTrainer base on filter input
//	@Tags			Trainer
//	@Accept			json
//	@Produce		json
//	@Param			filter_input	body		FilterTrainerInput	true	"put FilterTrainerInput input json and pass to  gin.Context"
//	@Success		200				{object}	responses.FilterTrainerResponses
//	@Security		BearerAuth
//	@Router			/protected/filter-trainer [post]
func FilterTrainer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input FilterTrainerInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.FilterTrainerResponses{
				Status:  http.StatusBadRequest,
				Message: "input missing",
			})
			return
		}
		fmt.Println("FilterTrainer input ", input)
		//--------
		result, err := models.FindFilteredTrainer(input.Specialty, input.Limit, input.FeeMin, input.FeeMax)
		fmt.Println(result)
		// result, err := models.FindProfile(input.Username, "trainer")
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.FilterTrainerResponses{
				Status:  http.StatusBadRequest,
				Message: `filter trainer profile  unsuccessful`,
			})
			return
		}
		c.JSON(http.StatusOK, responses.FilterTrainerResponses{
			Status:   http.StatusOK,
			Message:  `Successfully retrieve filtered trainer`,
			Trainers: result,
		})
		//-------------

		// if len(input.Specialty) == 0 {
		// 	fmt.Println("Etude")

		// } else {
		// 	result, err := models.FindFilteredTrainer(input.Specialty, input.Limit)
		// 	fmt.Println(result)
		// 	// result, err := models.FindProfile(input.Username, "trainer")
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		c.JSON(http.StatusBadRequest, responses.FilterTrainerResponse{
		// 			Status:  http.StatusBadRequest,
		// 			Message: `filter trainer profile  unsuccessful`,
		// 		})
		// 		return
		// 	}
		// 	c.JSON(http.StatusOK, responses.FilterTrainerResponses{
		// 		Status:   http.StatusOK,
		// 		Message:  `Successfully retrieve filtered trainer`,
		// 		Trainers: result,
		// 	})
		// }

	}
}
