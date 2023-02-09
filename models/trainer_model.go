package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindFilteredTrainer(speciality []string, limit int) ([]map[string]interface{}, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.TODO()
	// count := 0
	var results []map[string]interface{}
	var users []User
	// var err error
	// var cur *mongo.Cursor
	var opts *options.FindOptions
	var filter bson.D

	if len(speciality) == 0 {
		// fmt.Println("len speciality 0")
		// filter := bson.D{{"trainerInfo.speciality", bson.D{{"$in", speciality}}}}
		filter = bson.D{{Key: "usertype", Value: "Trainer"}}
		// opts = options.Find().SetLimit(int64(limit)).SetSort(bson.D{{"trainerInfo.rating", -1}, {"fee", 1}}).SetProjection(bson.D{{"_id", 0}, {"hashedPassword", 0}, {"createdAt", 0}, {"updatedAt", 0}})

	} else {
		filter = bson.D{{"trainerInfo.speciality", bson.D{{"$in", speciality}}}}
		filter = append(filter, bson.E{Key: "usertype", Value: "Trainer"})
		// opts = options.Find().SetLimit(int64(limit)).SetSort(bson.D{{"trainerInfo.rating", -1}, {"fee", 1}}).SetProjection(bson.D{{"_id", 0}, {"hashedPassword", 0}, {"createdAt", 0}, {"updatedAt", 0}})

	}
	opts = options.Find().SetLimit(int64(limit)).SetSort(bson.D{{"trainerInfo.rating", -1}, {"fee", 1}}).SetProjection(bson.D{{"_id", 0}, {"hashedPassword", 0}, {"createdAt", 0}, {"updatedAt", 0}, {"birthDate", 0}, {"citizenId", 0}, {"phoneNumber", 0}, {"userType", 0}})
	cur, err := userCollection.Find(ctx, filter, opts)

	if err != nil {
		// fmt.Println("error 1")
		return nil, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		// if count >= limit {
		// 	break
		// }
		var user User
		if err := cur.Decode(&user); err != nil {
			// fmt.Println("error 2")
			return nil, err
		}
		result := map[string]interface{}{
			// "usertype":    user.UserType,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
			// "birthdate":   user.BirthDate,
			// "citizenId":   user.CitizenId,
			"gender": user.Gender,
			// "phoneNumber": user.PhoneNumber,
			"address":     user.Address,
			"subAddress":  user.SubAddress,
			"avatarUrl":   user.AvatarUrl,
			"username":    user.Username,
			"trainerInfo": user.TrainerInfo,
		}
		// count += 1
		users = append(users, user)
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	// fmt.Println("print filter trainer", users)

	return results, nil
}

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
