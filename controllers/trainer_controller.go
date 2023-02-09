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
	Speciality []string `json:"Speciality" binding:"required"`
	// Rating     float32 `json:"Rating" binding:"required"`
	// Fee        float32 `json:"Fee" binding:"required"`
}
type TrainerInput struct {
	Speciality     []string `json:"speciality" `
	Rating         float64  `json:"rating "`
	Fee            float64  `json:"fee" `
	TraineeCount   int32    `json:"traineeCount" `
	CertificateUrl string   `json:"certificateUrl"`
}

type GetTrainerInput struct {
	Username string `json:"username" binding:"required"`
}

// FilterTrainer godoc
//
//	@Summary		FilterTrainer base on filter input
//	@Description	FilterTrainer base on filter input
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.FilterTrainerResponses
//
//	@Router			/protected/filter-trainer [get]
func FilterTrainer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input FilterTrainerInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.GetTrainerResponses{
				Status:  http.StatusBadRequest,
				Message: "input missing"})
			return
		}
		result, err := models.FindFilteredTrainer(input.Speciality)
		fmt.Println(result)
		// result, err := models.FindProfile(input.Username, "trainer")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
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
		result, err := models.FindProfile(input.Username, "Trainer")
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadRequest,
				Message: `trainer profile retrieval unsuccessful`,
			})
			return
		}
		c.JSON(http.StatusOK, responses.GetProfileResponses{
			Status:  http.StatusOK,
			Message: `Successfully retrieve trainer profile`,
			User:    result,
		})
	}
}

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
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		_, err = models.UpdateTrainerProfile(
			username, input.Speciality, input.Rating,
			input.Fee, input.TraineeCount, input.CertificateUrl)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadRequest,
				Message: `update failed`,
			})
			return
		}

		c.JSON(http.StatusOK,
			responses.ProfileResponses{
				Status:  http.StatusOK,
				Message: username + ` update success!`,
			})
	}
}
