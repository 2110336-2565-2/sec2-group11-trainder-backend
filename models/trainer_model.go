package models

import (
	"context"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FilteredTrainerInfo struct {
	Username    string      `json:"username"`
	FirstName   string      `json:"firstname"`
	LastName    string      `json:"lastname"`
	Gender      string      `json:"gender"`
	Address     string      `json:"address"`
	AvatarUrl   string      `json:"avatarUrl"`
	TrainerInfo TrainerInfo `json:"trainerInfo"`
}
type Review struct {
	Username  string    `bson:"username"`
	Rating    float64   `bson:"rating"`
	Comment   string    `bson:"comment"`
	CreatedAt time.Time `bson:"createdAt"`
}
type UserNotExist struct{}

func (e *UserNotExist) Error() string {
	return "error: user not existed"
}

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
		Lat:         user.Lat,
		Lng:         user.Lng,
	}
	return userProfile, user.TrainerInfo, nil
}

func userExists(username string) (bool, error) {
	filter := bson.M{"username": username}
	count, err := userCollection.CountDocuments(context.Background(), filter, nil)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func updateRatingByUsername(username string) error {
	projection := bson.M{"reviews": 1, "_id": 0}
	var user User
	err := userCollection.FindOne(context.Background(), bson.M{"username": username}, options.FindOne().SetProjection(projection)).Decode(&user)
	if err != nil {
		return err
	}
	var sum float64
	for _, review := range user.Reviews {
		sum += review.Rating
	}
	avgRating := math.Round(sum/float64(len(user.Reviews))*100) / 100
	_, err = userCollection.UpdateOne(context.Background(), bson.M{"username": username}, bson.M{"$set": bson.M{"trainerInfo.rating": avgRating}})
	if err != nil {
		return err
	}
	return nil
}

func AddReview(trainerUsername string, username string, rating float64, comment string) error {
	isExist, err := userExists(trainerUsername)
	if err != nil {
		return err
	}
	if !isExist {
		err = &UserNotExist{}
		return err
	}
	review := Review{
		Username:  username,
		Rating:    rating,
		Comment:   comment,
		CreatedAt: time.Now(),
	}
	filter := bson.M{"username": trainerUsername}
	update := bson.M{"$push": bson.M{"reviews": review}}
	_, err = userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	err = updateRatingByUsername(trainerUsername)
	if err != nil {
		return err
	}

	return nil

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

	// filter by fee
	if feeLowerBound == 0 && feeUpperBound == 0 {
		feeUpperBound = 1000000000
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
	update := bson.M{
		"$set": bson.M{
			"trainerInfo.specialty":      specialty,
			"updatedAt":                  time.Now(),
			"trainerInfo.rating":         rating,
			"trainerInfo.fee":            fee,
			"trainerInfo.traineeCount":   traineeCount,
			"trainerInfo.certificateUrl": certificateUrl,
		},
	}
	result, err = userCollection.UpdateOne(
		ctx,
		bson.M{"username": username},
		update,
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

func GetDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	const PI float64 = math.Pi
	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)
	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)
	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}
	dist = math.Acos(dist)

	return dist * 180 / PI * 1.609344 * 60 * 1.1515
}

func GetReviews(username string, limit int) ([]Review, error) {
	isExist, err := userExists(username)
	if err != nil {
		return nil, err
	}
	if !isExist {
		err = &UserNotExist{}
		return nil, err
	}
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pipeline := bson.A{
		bson.M{"$match": bson.M{"username": username}},
		bson.M{"$unwind": "$reviews"},
		bson.M{"$sort": bson.M{"reviews.createdAt": -1, "reviews.rating": -1}},
		bson.M{"$limit": limit},
		bson.M{"$group": bson.M{
			"_id":      "$_id",
			"username": bson.M{"$first": "$username"},
			"reviews":  bson.M{"$push": "$reviews"},
		}},
		bson.M{"$project": bson.M{
			"_id":      0,
			"username": 1,
			"reviews":  1,
		}},
	}
	limitOptions := options.Aggregate().SetMaxTime(2 * time.Second)

	var result struct {
		ReviewSlice []Review `bson:"reviews"`
	}
	// var reviews []Review
	cursor, err := userCollection.Aggregate(context.Background(), pipeline, limitOptions)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	defer cursor.Close(context.Background())
	// var reviews []Review
	reviews := make([]Review, 0)
	// Iterate through the results
	for cursor.Next(context.Background()) {
		// var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			// fmt.Println(err)
			return nil, err
		}
		// fmt.Println("result.MySlice", result.MySlice)

		for _, r := range result.ReviewSlice {
			// fmt.Println(r.Username)
			review := Review{
				Username:  r.Username,
				Rating:    r.Rating,
				Comment:   r.Comment,
				CreatedAt: r.CreatedAt,
			}
			reviews = append(reviews, review)
		}

	}

	return reviews, nil

}
