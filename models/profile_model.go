package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func FindProfile(username string) (userProfile bson.M, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result bson.M
	err = userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&result)

	return result, err

}
