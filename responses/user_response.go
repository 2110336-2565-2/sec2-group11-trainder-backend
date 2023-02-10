package responses

type CurrentUserResponse struct {
	Status   int    `json:"status"`
	Message  string `json:"message,omitempty"`
	Username string `json:"username,omitempty"`
}

type ProfileResponses struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type GetProfileResponses struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message,omitempty"`
	User    map[string]interface{} `json:"user,omitempty"`
}

type GetTrainerResponses struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message,omitempty"`
	User    map[string]interface{} `json:"user,omitempty"`
}

type FilterTrainerResponses struct {
	Status   int                      `json:"status"`
	Message  string                   `json:"message,omitempty"`
	Trainers []map[string]interface{} `json:"Trainers,omitempty"`
}
