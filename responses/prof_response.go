package responses

type ProfileResponses struct {
	Status  int    `jason:"status"`
	Message string `jason:"message,omitempty"`
}
