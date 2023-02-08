package controllers

import (
	"fmt"
	"net/http"

	"trainder-api/models"
	"trainder-api/responses"

	// "trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

type FilterTrainerInput struct {
	Speciality []string `json:"Speciality" binding:"required"`
	// Rating     float32 `json:"Rating" binding:"required"`
	// Fee        float32 `json:"Fee" binding:"required"`
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
			c.JSON(http.StatusBadRequest, responses.CurrentUserResponse{
				Status:  http.StatusBadRequest,
				Message: `filter trainer profile retrieval unsuccessful`,
			})
			return
		}
		c.JSON(http.StatusOK, responses.FilterTrainerResponses{
			Status:   http.StatusOK,
			Message:  `Successfully retrieve filtered trainer profile`,
			Trainers: result,
		})
	}
}
