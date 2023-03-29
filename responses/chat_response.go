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
	Status  int              `json:"status"`
	Message string           `json:"message,omitempty"`
	AllChat []models.AllChat `json:"allChat,omitempty"`
}

type PastChatResponse struct {
	Status       int              `json:"status"`
	Message      string           `json:"message,omitempty"`
	ChatMesseges []models.Messege `json:"chatMesseges,omitempty"`
}
