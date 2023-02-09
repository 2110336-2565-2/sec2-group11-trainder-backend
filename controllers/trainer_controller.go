package controllers

import (
	"net/http"

	"trainder-api/models"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

type TrainerInput struct {
	Speciality     []string `json:"speciality"`
	Rating         float64  `json:"rating"`
	Fee            float64  `json:"fee"`
	TraineeCount   int32    `json:"traineeCount"`
	CertificateUrl string   `json:"certificateUrl"`
}

type GetTrainerInput struct {
	Username string `json:"username" binding:"required"`
}

// GetTrainer retrieves the trainer profile of the user who made the request
// @Summary	Retrieve trainer profile
// @Description	Retrieves the trainer profile information of the user who made the request.
// @Tags Trainer
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param input body GetTrainerInput true "Put username input for retrieving the trainer profile"
// @Success 200 {object} responses.GetTrainerResponses
// @Failure 400 {object} responses.GetTrainerResponses
// @Router /protected/trainer [post]
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
				Message: `trainer profile retrieval unsuccessful`,
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
// @Summary Update the trainer's profile information.
// @Tags Trainer
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile body TrainerInput true "Trainer's information to update"
// @Success 200 {object} responses.ProfileResponses
// @Failure 400 {object} responses.ProfileResponses
// @Failure 401 {object} responses.ProfileResponses
// @Router /protected/update-trainer [post]
func UpdateTrainer() gin.HandlerFunc {

	return func(c *gin.Context) {
		var input TrainerInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponses{
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
			username, input.Speciality, input.Rating,
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
