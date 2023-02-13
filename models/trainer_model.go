package models

import (
	"context"
	// "fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Use for only finding the profile of a trainer, which will have normal user profile and trainer info
func FindTrainerProfile(username string) (userProfile UserProfile, trainerInfo TrainerInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "username", Value: username}, {Key: "usertype", Value: "Trainer"}}
	opts := options.FindOne().SetProjection(bson.D{
		{Key: "_id", Value: 0},
		{Key: "hashedPassword", Value: 0},
		{Key: "createdAt", Value: 0},
		{Key: "updatedAt", Value: 0}})
	var user User
	err = userCollection.FindOne(ctx, filter, opts).Decode(&user)
	if err != nil {
		return userProfile, trainerInfo, err
	}
	userProfile = UserProfile{
		Username:    user.Username,
		UserType:    user.UserType,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		BirthDate:   user.BirthDate.Format("2006-01-02"),
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		AvatarUrl:   user.AvatarUrl,
	}
	return userProfile, user.TrainerInfo, nil
}

type FilteredTrainerInfo struct {
	Username    string      `json:"username"`
	FirstName   string      `json:"firstname"`
	LastName    string      `json:"lastname"`
	Gender      string      `json:"gender"`
	Address     string      `json:"address"`
	AvatarUrl   string      `json:"avatarUrl"`
	TrainerInfo TrainerInfo `json:"trainerInfo"`
}

func FindFilteredTrainer(specialty []string, limit int, feeLowerBound float64, feeUpperBound float64) ([]FilteredTrainerInfo, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.TODO()

	var results []FilteredTrainerInfo
	var opts *options.FindOptions
	var filterArr []bson.D

	if len(specialty) == 0 {
		// filter := bson.D{{"trainerInfo.specialty", bson.D{{"$in", specialty}}}}
		filterArr = append(filterArr, bson.D{{Key: "usertype", Value: "Trainer"}})
		// opts = options.Find().SetLimit(int64(limit)).SetSort(bson.D{{"trainerInfo.rating", -1}, {"fee", 1}}).SetProjection(bson.D{{"_id", 0}, {"hashedPassword", 0}, {"createdAt", 0}, {"updatedAt", 0}})

	} else {
		filterArr = append(filterArr, bson.D{{Key: "trainerInfo.specialty", Value: bson.D{{Key: "$in", Value: specialty}}}})
		filterArr = append(filterArr, bson.D{{Key: "usertype", Value: "Trainer"}})
		// opts = options.Find().SetLimit(int64(limit)).SetSort(bson.D{{"trainerInfo.rating", -1}, {"fee", 1}}).SetProjection(bson.D{{"_id", 0}, {"hashedPassword", 0}, {"createdAt", 0}, {"updatedAt", 0}})

	}

	// ------------------------filter by fee----------------------------
	if feeLowerBound == 0 && feeUpperBound == 0 {
		feeUpperBound = 1000000000 //intmax
	}
	filter2 := bson.D{
		{
			Key: "trainerInfo.fee",
			Value: bson.M{
				"$gte": feeLowerBound,
				"$lte": feeUpperBound,
			},
		},
	}

	filterArr = append(filterArr, filter2)

	filter := bson.M{
		"$and": filterArr,
	}

	opts = options.Find().SetLimit(int64(limit)).SetSort(bson.D{
		{Key: "trainerInfo.rating", Value: -1},
		{Key: "fee", Value: 1},
	}).SetProjection(bson.D{
		{Key: "_id", Value: 0},
		{Key: "hashedPassword", Value: 0},
		{Key: "createdAt", Value: 0},
		{Key: "updatedAt", Value: 0},
		{Key: "birthDate", Value: 0},
		{Key: "citizenId", Value: 0},
		{Key: "phoneNumber", Value: 0},
		{Key: "userType", Value: 0}})
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
		result := FilteredTrainerInfo{
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Gender:      user.Gender,
			Address:     user.Address,
			AvatarUrl:   user.AvatarUrl,
			Username:    user.Username,
			TrainerInfo: user.TrainerInfo,
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func UpdateTrainerProfile(username string, specialty []string, rating float64, fee float64, traineeCount int32, certificateUrl string) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	update := bson.M{}
	if len(specialty) > 0 {
		update["specialty"] = specialty
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
