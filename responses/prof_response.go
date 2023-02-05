package responses

import (
	"go.mongodb.org/mongo-driver/bson"
)

type ProfileResponses struct {
	Status  int    `jason:"status"`
	Message string `jason:"message,omitempty"`
}

type GetProfileResponses struct {
	Status       int    `jason:"status"`
	Profile_info bson.M `bson:"message,omitempty"`
}
