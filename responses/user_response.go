package responses

type CurrentUserResponse struct {
	Status   int    `json:"status"`
	Message  string `json:"message,omitempty"`
	Username string `json:"username,omitempty"`
}
