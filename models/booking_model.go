package models

import (
	"context"
	"fmt"
	"time"
	"trainder-api/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var bookingsCollection *mongo.Collection = configs.GetCollection(configs.DB, "bookings")

type Booking struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	Trainer       string             `bson:"trainer" json:"trainer"`
	Trainee       string             `bson:"trainee" json:"trainee"`
	StartDateTime time.Time          `bson:"startDateTime" json:"startDateTime"`
	EndDateTime   time.Time          `bson:"endDateTime" json:"endDateTime"`
	Status        string             `bson:"status" json:"status"`
	Payment       struct {
		TotalCost float64 `bson:"totalCost" json:"totalCost"`
		Status    string  `bson:"status" json:"status"`
	} `bson:"payment" json:"payment"`
}

func CreateBooking(trainee string, trainer string, date string, startTime string, endTime string) error {
	// fmt.Println(trainer)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"username": trainer, "usertype": "Trainer"}
	projection := bson.M{"trainerInfo.fee": 1, "_id": 0}
	var result struct {
		TrainerInfo struct {
			Fee float64 `bson:"fee"`
		} `bson:"trainerInfo"`
	}
	if err := userCollection.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&result); err != nil {
		return fmt.Errorf("failed to get trainer info: %w", err)
	}

	startDateTimeStr := date + " " + startTime
	startDateTime, err := time.Parse("2006-01-02 15:04", startDateTimeStr)
	if err != nil {
		return fmt.Errorf("failed to parse start datetime: %v", err)
	}
	endDateTimeStr := date + " " + endTime
	endDateTime, err := time.Parse("2006-01-02 15:04", endDateTimeStr)
	if err != nil {

		return fmt.Errorf("failed to parse end datetime: %v", err)
	}
	duration := endDateTime.Sub(startDateTime)
	totalCost := result.TrainerInfo.Fee * duration.Hours()

	// Create booking object
	booking := bson.M{
		"trainer":       trainer,
		"trainee":       trainee,
		"startDateTime": startDateTime,
		"endDateTime":   endDateTime,
		"status":        "pending",
		"payment": bson.M{
			"totalCost": totalCost,
			"status":    "pending",
		},
	}

	// Insert booking into database
	_, err = bookingsCollection.InsertOne(ctx, booking)
	if err != nil {
		return fmt.Errorf("failed to create booking: %v", err)
	}

	return nil
}

// merge into one function ()
func GetUpcomingBookings(Username string) ([]Booking, error) {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var filter bson.M
	if IsTrainer(Username) {
		filter = bson.M{
			"trainer": Username,
			"startDateTime": bson.M{
				"$gt": now,
			},
		}
	} else {
		filter = bson.M{
			"trainee": Username,
			"startDateTime": bson.M{
				"$gt": now,
			},
		}
	}

	cursor, err := bookingsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// fmt.Println(cursor)
	var bookings []Booking
	for cursor.Next(ctx) {
		var booking Booking
		err := cursor.Decode(&booking)
		if err != nil {
			return nil, err
		}
		// fmt.Println(booking)
		bookings = append(bookings, booking)
	}

	return bookings, nil

}

// func GetUpcomingBookingsForTrainee(trainerUsername string) ([]Booking, error) {
// 	now := time.Now()
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	filter := bson.M{
// 		"trainee": trainerUsername,
// 		"startDateTime": bson.M{
// 			"$gt": now,
// 		},
// 	}
// 	cursor, err := bookingsCollection.Find(ctx, filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// fmt.Println(cursor)
// 	var bookings []Booking
// 	for cursor.Next(ctx) {
// 		var booking Booking
// 		err := cursor.Decode(&booking)
// 		if err != nil {
// 			return nil, err
// 		}
// 		// fmt.Println(booking)
// 		bookings = append(bookings, booking)
// 	}

// 	return bookings, nil

// }
// func GetUpcomingBookingsForTrainer(trainerUsername string) ([]Booking, error) {
// 	now := time.Now()
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	filter := bson.M{
// 		"trainer": trainerUsername,
// 		"startDateTime": bson.M{
// 			"$gt": now,
// 		},
// 	}
// 	cursor, err := bookingsCollection.Find(ctx, filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// fmt.Println(cursor)
// 	var bookings []Booking
// 	for cursor.Next(ctx) {
// 		var booking Booking
// 		err := cursor.Decode(&booking)
// 		if err != nil {
// 			return nil, err
// 		}
// 		// fmt.Println(booking)
// 		bookings = append(bookings, booking)
// 	}

// 	return bookings, nil

// }

func UpdateBooking(bookingObjectId string, status string, paymentStatus string) error {
	objectID, err := primitive.ObjectIDFromHex(bookingObjectId)
	if err != nil {
		return fmt.Errorf("failed to parse bookingObjId: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateArr := bson.M{}

	if len(status) != 0 {
		updateArr["status"] = status

	}
	if len(paymentStatus) != 0 {
		updateArr["payment.status"] = paymentStatus
	}

	res, err := bookingsCollection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": updateArr},
	)
	if err != nil {
		return fmt.Errorf("failed to UpdateOne bookingsCollection: %v", err)
	}
	// fmt.Println("res", res, res.MatchedCount)
	if res.MatchedCount == 0 {
		return fmt.Errorf("the bookingObjectId could not be found")
	}

	return nil
}

func DeleteBooking(bookingObjectId string) error {
	objectID, err := primitive.ObjectIDFromHex(bookingObjectId)
	if err != nil {
		return fmt.Errorf("failed to parse bookingObjId: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := bookingsCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("failed to delete document on bookingsCollection: %v", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("the bookingObjectId could not be found")
	}
	return nil
}
