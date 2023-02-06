package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindProfile(username string) (userProfile bson.M, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result bson.M
	err = userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&result)

	return result, err

}

func UpdateUserProfile(username string, firstName string, lastName string, birthDate string, citizenID string, gender string, phoneNumber string, address string, subAddress string) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	layout := "2006-01-02"
	date, error := time.Parse(layout, birthDate)

	if error != nil {
		return
	}

	result, err = userCollection.UpdateOne(
		ctx,
		bson.M{"username": username},
		bson.M{"$set": bson.M{
			"firstname":   firstName,
			"lastname":    lastName,
			"birthdate":   date,
			"citizenId":   citizenID,
			"gender":      gender,
			"phoneNumber": phoneNumber,
			"address":     address,
			"subAddress":  subAddress,
			"updatedAt":    time.Now(),
		}},
	)
	return
}
