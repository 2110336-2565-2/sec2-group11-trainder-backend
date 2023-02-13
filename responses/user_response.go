package responses

import (
	"trainder-api/models"
)

type CurrentUserResponse struct {
	Status   int    `json:"status"`
	Message  string `json:"message,omitempty"`
	Username string `json:"username,omitempty"`
}

type ProfileResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type UserProfileResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message,omitempty"`
	User    models.UserProfile `json:"user,omitempty"`
}

type TrainerProfileResponse struct {
	Status      int                `json:"status"`
	Message     string             `json:"message,omitempty"`
	User        models.UserProfile `json:"user,omitempty"`
	TrainerInfo models.TrainerInfo `json:"trainerInfo,omitempty"`
}

type FilterTrainerResponses struct {
	Status   int                          `json:"status"`
	Message  string                       `json:"message,omitempty"`
	Trainers []models.FilteredTrainerInfo `json:"trainers,omitempty"`
}
