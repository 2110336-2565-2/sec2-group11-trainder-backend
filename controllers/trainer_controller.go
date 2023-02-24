package controllers

import (
	"net/http"

	"trainder-api/models"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

type FilterTrainerForm struct {
	Specialty []string `json:"specialty"`
	Limit     int      `json:"limit" binding:"required"`
	FeeMin    float64  `json:"feeMin"`
	FeeMax    float64  `json:"feeMax"`
}
type UpdateTrainerDetails struct {
	Specialty      []string `json:"specialty"`
	Rating         float64  `json:"rating"`
	Fee            float64  `json:"fee"`
	TraineeCount   int32    `json:"traineeCount"`
	CertificateUrl string   `json:"certificateUrl"`
}
type GetTrainerForm struct {
	Username string `json:"username" binding:"required"`
}

type ReviewDetails struct {
	TrainerUsername string  `json:"trainerUsername" binding:"required"`
	Rating          float64 `json:"rating" binding:"required"`
	Comment         string  `json:"comment" binding:"required"`
}

type GetReviewsForm struct {
	TrainerUsername string `json:"trainerUsername" binding:"required"`
	Limit           int    `json:"limit" binding:"required"`
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

// GetTrainerProfile retrieves the trainer profile of any trainer
//
//	@Summary		Retrieve trainer profile
//	@Description	Retrieves the trainer profile information.
//	@Tags			Trainer
//	@Accept			json
//	@Produce		json
//	@Param			input	body		GetTrainerForm true					"Put username input for retrieving the trainer profile"
//	@Success		200		{object}	responses.TrainerProfileResponse	"Successfully retrieved the trainer profile"
//	@Failure		400		{object}	responses.TrainerProfileResponse	"Failed to retrieve the trainer profile"
//	@Security		BearerAuth
//	@Router			/protected/trainer [post]
func GetTrainerProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input GetTrainerForm
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
//	@Param		profile	body		UpdateTrainerDetails true	"Trainer's information to update"
//	@Success	200		{object}	responses.ProfileResponse	"Successfully update the trainer's profile"
//	@Failure	400		{object}	responses.ProfileResponse	"Bad Request, either invalid input or user is not a trainer"
//	@Failure	401		{object}	responses.ProfileResponse	"Unauthorized, the user is not logged in"
//	@Security	BearerAuth
//	@Router		/protected/update-trainer [post]
func UpdateTrainerProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input UpdateTrainerDetails
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
//	@Param			FilterTrainer	body		FilterTrainerForm true	"Parameters for filtering trainers"
//	@Success		200				{object}	responses.FilterTrainerResponse
//	@Failure		400				{object}	responses.FilterTrainerResponse
//	@Security		BearerAuth
//	@Router			/protected/filter-trainer [post]
func FilterTrainer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input FilterTrainerForm
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.FilterTrainerResponse{
				Status:  http.StatusBadRequest,
				Message: "input missing",
			})
			return
		}
		result, err := models.FindFilteredTrainer(input.Specialty, input.Limit, input.FeeMin, input.FeeMax)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.FilterTrainerResponse{
				Status:  http.StatusBadRequest,
				Message: `filter trainer profile unsuccessful`,
			})
			return
		}
		c.JSON(http.StatusOK, responses.FilterTrainerResponse{
			Status:   http.StatusOK,
			Message:  `Successfully retrieve filtered trainer`,
			Trainers: result,
		})

		// if len(input.Specialty) == 0 {

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

// @Summary		Get reviews of specific trainer
// @Description	Get reviews of specific trainer username from database sort by recent date then rating desc, limit number of output by limit
// @Tags		Trainer
// @Accept		json
// @Produce		json
// @Param		GetReviewsInput	body		GetReviewsForm true	"Parameters for querying trainer reviews"
// @Success		200				{object}	responses.TrainerReviewsResponse
// @Failure		400				{object}	responses.TrainerReviewsResponse
// @Security	BearerAuth
// @Router		/protected/get-reviews [post]
func GetReviews() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input GetReviewsForm

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainerReviewsResponse{
				Status:  http.StatusBadRequest,
				Message: "input missing",
			})
			return
		}
		// fmt.Println("input", input)
		result, err := models.GetReviews(input.TrainerUsername, input.Limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainerReviewsResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, responses.TrainerReviewsResponse{
			Status:  http.StatusOK,
			Message: `Successfully retrieve reviews of trainer` + input.TrainerUsername,
			Reviews: result,
		})
	}
}

// @Summary		Add trainer review
// @Description	Add review on trainer to database
// @Tags		Trainer
// @Accept		json
// @Produce		json
// @Param		ReviewRequest	body		ReviewDetails	true	"Parameters for trainer review"
// @Success		200				{object}	responses.AddReviewResponse
// @Failure		400				{object}	responses.AddReviewResponse
// @Security	BearerAuth
// @Router		/protected/add-review [post]
func AddTrainerReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ReviewDetails
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.AddReviewResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.AddReviewResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		err = models.AddReview(input.TrainerUsername, username, input.Rating, input.Comment)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.AddReviewResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,
			responses.RegisterResponse{
				Status:  http.StatusOK,
				Message: input.TrainerUsername + ` update success!`,
			})
	}
}
