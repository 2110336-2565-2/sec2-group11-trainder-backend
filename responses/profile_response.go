package responses

import (
	"go.mongodb.org/mongo-driver/bson"
)

type ProfileResponses struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type GetProfileResponses struct {
	Status      int    `json:"status"`
	ProfileInfo bson.M `bson:"message,omitempty"`
}
