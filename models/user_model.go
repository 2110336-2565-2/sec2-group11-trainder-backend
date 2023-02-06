package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindProfile(username string) (resultMap map[string]interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result bson.M
	err = userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&result)
	if err != nil {
		return nil, err
	}

	delete(result, "_id")
	delete(result, "createdAt")
	delete(result, "updatedAt")
	delete(result, "hashedPassword")
	resultMap = result
	for k, v := range resultMap {
		fmt.Println(k, "value is", v)
	}
	return resultMap, nil

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
