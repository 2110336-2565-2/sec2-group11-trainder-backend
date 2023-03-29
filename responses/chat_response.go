package responses

import "trainder-api/models"

type ChatResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}
type ChatRoomIDResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	RoomID  string `json:"roomID,omitempty"`
}

type AllChatResponse struct {
	Status   int              `json:"status"`
	Message  string           `json:"message,omitempty"`
	Response []models.AllChat `json:"roomID,omitempty"`
}
