package responses
type ChatResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}