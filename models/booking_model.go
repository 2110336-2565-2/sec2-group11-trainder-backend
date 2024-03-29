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

type Payment struct {
	TotalCost     int64  `bson:"totalCost" json:"totalCost"`
	Status        string `bson:"status" json:"status"`
	ChargeID      string `bson:"chargeID" json:"chargeID"`
	Bank          string `bson:"bank" json:"bank"`
	AccountName   string `bson:"accountName" json:"accountName"`
	AccountNumber string `bson:"accountNumber" json:"accountNumber"`
}

type Booking struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	Trainer       string             `bson:"trainer" json:"trainer"`
	Trainee       string             `bson:"trainee" json:"trainee"`
	StartDateTime time.Time          `bson:"startDateTime" json:"startDateTime"`
	EndDateTime   time.Time          `bson:"endDateTime" json:"endDateTime"`
	Status        string             `bson:"status" json:"status"`
	Payment       Payment            `bson:"payment" json:"payment"`
}

type ReturnBooking struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Trainer          string             `bson:"trainer" json:"trainer"`
	TrainerFirstName string             `bson:"trainerFirstName" json:"trainerFirstName"`
	TrainerLastName  string             `bson:"trainerLastName" json:"trainerLastName"`
	Trainee          string             `bson:"trainee" json:"trainee"`
	TraineeFirstName string             `bson:"traineeFirstName" json:"traineeFirstName"`
	TraineeLastName  string             `bson:"traineeLastName" json:"traineeLastName"`
	StartDateTime    time.Time          `bson:"startDateTime" json:"startDateTime"`
	EndDateTime      time.Time          `bson:"endDateTime" json:"endDateTime"`
	Status           string             `bson:"status" json:"status"`
	Payment          Payment            `bson:"payment" json:"payment"`
}

func CreateBooking(trainee string, trainer string, date string, startTime string, endTime string) error {
	// fmt.Println(trainer)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"username": trainer, "usertype": "Trainer"}
	projection := bson.M{"trainerInfo.fee": 1, "_id": 0}
	var result struct {
		TrainerInfo struct {
			Fee int64 `bson:"fee"`
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
	totalCost := result.TrainerInfo.Fee * (int64(duration.Hours()) + 1)

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
			"chargeID":  "",
		},
	}

	// Insert booking into database
	_, err = bookingsCollection.InsertOne(ctx, booking)
	if err != nil {
		return fmt.Errorf("failed to create booking: %v", err)
	}

	return nil
}

func GetBooking(bookingID string) (result Booking, err error) {
	objectID, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return result, err
	}

	filter := bson.M{"_id": objectID}

	err = bookingsCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, err
}

// merge into one function ()
func GetUpcomingBookings(Username string) ([]ReturnBooking, error) {
	now := time.Now().Local()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var filter bson.M
	if IsTrainer(Username) {
		filter = bson.M{
			"trainer": Username,
			"startDateTime": bson.M{
				"$gte": now.UTC().Truncate(24 * time.Hour),
			},
		}
	} else {
		filter = bson.M{
			"trainee": Username,
			"startDateTime": bson.M{
				"$gte": now.UTC().Truncate(24 * time.Hour),
			},
		}
	}

	cursor, err := bookingsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var bookings []ReturnBooking
	for cursor.Next(ctx) {
		var booking Booking
		err := cursor.Decode(&booking)
		if err != nil {
			return nil, err
		}
		var trainerInfo User
		err = userCollection.FindOne(ctx, bson.M{"username": booking.Trainer}).Decode(&trainerInfo)
		if err != nil {
			return nil, err
		}

		var traineeInfo User
		err = userCollection.FindOne(ctx, bson.M{"username": booking.Trainee}).Decode(&traineeInfo)
		if err != nil {
			return nil, err
		}
		result := ReturnBooking{
			ID:               booking.ID,
			Trainer:          booking.Trainer,
			TrainerFirstName: trainerInfo.FirstName,
			TrainerLastName:  trainerInfo.LastName,
			Trainee:          booking.Trainee,
			TraineeFirstName: traineeInfo.FirstName,
			TraineeLastName:  traineeInfo.LastName,
			StartDateTime:    booking.StartDateTime,
			EndDateTime:      booking.EndDateTime,
			Status:           booking.Status,
			Payment:          booking.Payment,
		}

		bookings = append(bookings, result)
	}

	return bookings, nil

}

func UpdateBooking(bookingObjectId string, status string, username string) error {
	booking, err := GetBooking(bookingObjectId)
	if err != nil {
		return err
	}

	if username != booking.Trainer {
		return fmt.Errorf("can only update own booking")
	}

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
	filter := bson.M{"_id": objectID}
	var bookingDoc ReturnBooking
	err = bookingsCollection.FindOne(ctx, filter).Decode(&bookingDoc)
	if err != nil {
		return fmt.Errorf("couldn't find bookingDoc")
	}

	// Only check time slot when confirming booking
	if status == "confirm" {
		startTime := bookingDoc.StartDateTime
		endTime := bookingDoc.EndDateTime
		filter = bson.M{
			"trainer": username,
			"status":  bson.M{"$in": []string{"confirm", "complete"}},
			"$or": []bson.M{
				{"$and": []bson.M{
					{"startDateTime": bson.M{"$gte": startTime}},
					{"startDateTime": bson.M{"$lt": endTime}},
				}},
				{"$and": []bson.M{
					{"endDateTime": bson.M{"$gt": startTime}},
					{"endDateTime": bson.M{"$lte": endTime}},
				}},
			},
		}
		count, err := bookingsCollection.CountDocuments(ctx, filter)
		if err != nil {
			return fmt.Errorf("couldn't update status")
		}
		if count > 0 {
			return fmt.Errorf("trainer has already confirmed another booking")
		}
	}

	res, err := bookingsCollection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": updateArr},
	)
	if err != nil {
		return fmt.Errorf("failed to UpdateOne bookingsCollection: %v", err)
	}
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

func GetSpecificDayBookings(Username string, date string) ([]ReturnBooking, error) {
	startDateTimeStr := date + " " + "00:00"
	datetime, err := time.Parse("2006-01-02 15:04", startDateTimeStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start datetime: %v", err)
	}
	today := datetime.Truncate(24 * time.Hour)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var filter bson.M
	if IsTrainer(Username) {
		filter = bson.M{
			"trainer": Username,
			"$expr": bson.M{
				"$eq": bson.A{
					bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$startDateTime"}},
					today.Format("2006-01-02"),
				},
			},
		}
	} else {
		filter = bson.M{
			"trainee": Username,
			"$expr": bson.M{
				"$eq": bson.A{
					bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$startDateTime"}},
					today.Format("2006-01-02"),
				},
			},
		}
	}
	cursor, err := bookingsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var bookings []ReturnBooking
	for cursor.Next(ctx) {
		var booking Booking
		err := cursor.Decode(&booking)
		if err != nil {
			return nil, err
		}
		var trainerInfo User
		err = userCollection.FindOne(ctx, bson.M{"username": booking.Trainer}).Decode(&trainerInfo)
		if err != nil {
			return nil, err
		}

		var traineeInfo User
		err = userCollection.FindOne(ctx, bson.M{"username": booking.Trainee}).Decode(&traineeInfo)
		if err != nil {
			return nil, err
		}
		result := ReturnBooking{
			ID:               booking.ID,
			Trainer:          booking.Trainer,
			TrainerFirstName: trainerInfo.FirstName,
			TrainerLastName:  trainerInfo.LastName,
			Trainee:          booking.Trainee,
			TraineeFirstName: traineeInfo.FirstName,
			TraineeLastName:  traineeInfo.LastName,
			StartDateTime:    booking.StartDateTime,
			EndDateTime:      booking.EndDateTime,
			Status:           booking.Status,
			Payment:          booking.Payment,
		}
		bookings = append(bookings, result)
	}

	return bookings, nil
}
