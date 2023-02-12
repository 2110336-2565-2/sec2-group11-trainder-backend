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
	Lower_Fee float64  `json:"Lower_Fee"`
	Upper_Fee float64  `json:"Upper_Fee"`
	// Rating     float64 `json:"Rating" binding:"required"`
}
type TrainerInput struct {
	Specialty      []string `json:"specialty"`
	Rating         float64  `json:"rating"`
	Fee            float64  `json:"fee"`
	TraineeCount   int32    `json:"traineeCount"`
	CertificateUrl string   `json:"certificateUrl"`
}

type GetTrainerInput struct {
	Username string `json:"username" binding:"required"`
}

// FilterTrainer godoc
//
//	@Summary		FilterTrainer base on filter input
//	@Description	FilterTrainer base on filter input
//	@Tags			Trainer
//	@Accept			json
//	@Produce		json
//	@Param			filter_input	body		FilterTrainerInput	true	"put FilterTrainerInput input json and pass to  gin.Context"
//	@Success		200					{object}	responses.FilterTrainerResponses
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
		result, err := models.FindFilteredTrainer(input.Specialty, input.Limit, input.Lower_Fee, input.Upper_Fee)
		fmt.Println(result)
		// result, err := models.FindProfile(input.Username, "trainer")
		if err != nil {
			fmt.Println(err)
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

// GetTrainer retrieves the trainer profile of the user who made the request
//
//	@Summary		Retrieve trainer profile
//	@Description	Retrieves the trainer profile information of the user who made the request.
//	@Tags			Trainer
//	@Accept			json
//	@Produce		json
//	@Param			input	body		GetTrainerInput					true	"Put username input for retrieving the trainer profile"
//	@Success		200		{object}	responses.GetProfileResponses	"Successfully retrieved the trainer profile"
//	@Failure		400		{object}	responses.GetTrainerResponses	"Failed to retrieve the trainer profile"
//	@Security		BearerAuth
//	@Router			/protected/trainer [post]
func GetTrainer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input GetTrainerInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.GetTrainerResponses{
				Status:  http.StatusBadRequest,
				Message: "input missing"})
			return
		}
		result, err := models.FindProfile(input.Username, "Trainer")
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GetTrainerResponses{
				Status:  http.StatusBadRequest,
				Message: "Failed to retrieve the trainer profile",
			})
			return
		}
		c.JSON(http.StatusOK, responses.GetTrainerResponses{
			Status:  http.StatusOK,
			Message: `Successfully retrieve trainer profile`,
			User:    result,
		})
	}
}

// Update the trainer's profile information.
//
//	@Summary	Update the trainer's profile information.
//	@Tags		Trainer
//	@Accept		json
//	@Produce	json
//	@Param		profile	body		TrainerInput				true	"Trainer's information to update"
//	@Success	200		{object}	responses.ProfileResponses	"Successfully update the trainer's profile"
//	@Failure	400		{object}	responses.ProfileResponses	"Bad Request, either invalid input or user is not a trainer"
//	@Failure	401		{object}	responses.ProfileResponses	"Unauthorized, the user is not logged in"
//	@Security	BearerAuth
//	@Router		/protected/update-trainer [post]
func UpdateTrainer() gin.HandlerFunc {

	return func(c *gin.Context) {
		var input TrainerInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.GetTrainerResponses{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil || !models.IsTrainer(username) {
			c.JSON(http.StatusBadRequest, responses.GetTrainerResponses{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		_, err = models.UpdateTrainerProfile(
			username, input.Specialty, input.Rating,
			input.Fee, input.TraineeCount, input.CertificateUrl)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GetTrainerResponses{
				Status:  http.StatusBadRequest,
				Message: `update failed`,
			})
			return
		}

		c.JSON(http.StatusOK,
			responses.GetTrainerResponses{
				Status:  http.StatusOK,
				Message: username + ` update success!`,
			})
	}
}
