package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindFilteredTrainer(speciality []string) ([]map[string]interface{}, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.TODO()
	// defer cancel()
	filter := bson.D{{"trainerInfo.speciality", bson.D{{"$in", speciality}}}}
	filter = append(filter, bson.E{Key: "usertype", Value: "Trainer"})
	opts := options.Find().SetSort(bson.M{"name": 1}).SetProjection(bson.D{{"_id", 0}, {"hashedPassword", 0}, {"createdAt", 0}, {"updatedAt", 0}})
	var results []map[string]interface{}
	// var users []User
	cur, err := userCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		result := map[string]interface{}{
			"usertype":    user.UserType,
			"firstname":   user.FirstName,
			"lastname":    user.LastName,
			"birthdate":   user.BirthDate,
			"citizenId":   user.CitizenId,
			"gender":      user.Gender,
			"phoneNumber": user.PhoneNumber,
			"address":     user.Address,
			"subAddress":  user.SubAddress,
			"avatarUrl":   user.AvatarUrl,
			"username":    user.Username,
			"trainerInfo": user.TrainerInfo,
		}
		// users = append(users, user)
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	// fmt.Println("print filter trainer", users)

	return results, nil
}
