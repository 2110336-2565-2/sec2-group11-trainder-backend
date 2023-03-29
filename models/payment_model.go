package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentInfo struct {
	TraineeUsername string
	TotalCost       int64
	Status          string
}

func GetPaymentInfo(bookingID string) (PaymentInfo, error) {
	// Filter using booking id
	objectID, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return PaymentInfo{}, err
	}

	filter := bson.M{"_id": objectID}

	var result Booking

	err = bookingsCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return PaymentInfo{}, fmt.Errorf("no booking matched the requested id")
		}
		return PaymentInfo{}, err
	}

	return PaymentInfo{
		TraineeUsername: result.Trainee,
		TotalCost:       int64(result.Payment.TotalCost),
		Status:          result.Payment.Status,
	}, err
}

func Pay(bookingID string, chargeId string) (err error) {
	objectID, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateStatus, err := bookingsCollection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"status": "complete", "payment.status": "paid", "payment.chargeID": chargeId}})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("no booking matched the requested id")
		}
		return err
	}
	if updateStatus.MatchedCount == 0 {
		return fmt.Errorf("the bookingObjectId could not be found")
	}

	return err

}
