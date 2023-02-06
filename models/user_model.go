package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindProfile(username string) (result map[string]interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "username", Value: username}}
	opts := options.Find().SetProjection(bson.D{{"_id", 0}, {"hashedPassword", 0}, {"createdAt", 0}, {"updatedAt", 0}})
	cursor, err := userCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	var user []User
	for cursor.Next(ctx) {
		var u User
		if err := cursor.Decode(&u); err != nil {
			return nil, err
		}
		user = append(user, u)
	}
	result = map[string]interface{}{
		"firstname":   user[0].FirstName,
		"lastname":    user[0].LastName,
		"birthdate":   user[0].BirthDate,
		"citizenId":   user[0].CitizenId,
		"gender":      user[0].Gender,
		"phoneNumber": user[0].PhoneNumber,
		"address":     user[0].Address,
		"subAddress":  user[0].SubAddress,
	}

	return result, nil

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
			"updatedAt":   time.Now(),
		}},
	)
	return
}
