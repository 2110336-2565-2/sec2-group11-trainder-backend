package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateTrainerProfile(username string, speciality []string, rating float64, fee float64, traineeCount int32, certificateUrl string) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	update := bson.M{}
	if len(speciality) > 0 {
		update["speciality"] = speciality
	}
	if rating > 0 {
		update["rating"] = rating
	}
	if fee > 0 {
		update["fee"] = fee
	}
	if traineeCount > 0 {
		update["traineeCount"] = traineeCount
	}
	if certificateUrl != "" {
		update["certificateUrl"] = certificateUrl
	}
	result, err = userCollection.UpdateOne(
		ctx,
		bson.M{"username": username},
		bson.M{"$set": bson.M{
			"trainerInfo": update,
			"updatedAt":   time.Now(),
		}},
	)
	return
}

func IsTrainer(username string) (b bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "username", Value: username}}
	var user User
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil || user.UserType != "Trainer" {
		return false
	}
	return true
}
