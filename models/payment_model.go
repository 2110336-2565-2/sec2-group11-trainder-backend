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
	TrainerUsername string
	TotalCost       int64
	BookingStatus   string
	PaymentStatus   string
}

func GetPaymentInfo(bookingID string) (PaymentInfo, error) {
	booking, err := GetBooking(bookingID)

	return PaymentInfo{
		TraineeUsername: booking.Trainee,
		TrainerUsername: booking.Trainer,
		TotalCost:       int64(booking.Payment.TotalCost),
		BookingStatus:   booking.Status,
		PaymentStatus:   booking.Payment.Status,
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

	updateStatus, err := bookingsCollection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"payment.status": "paid", "payment.chargeID": chargeId}})

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

func RequestPayout(bookingID string) (err error) {
	objectID, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateStatus, err := bookingsCollection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"payment.status": "need_payout"}})

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

func Payout(bookingID string) (err error) {
	objectID, err := primitive.ObjectIDFromHex(bookingID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateStatus, err := bookingsCollection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"payment.status": "paid_out"}})

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

func PaymentNeedPayouts() (payments []Payment, err error) {
	filter := bson.M{"payment.status": "need_payout"}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var results []Booking

	cursor, err := bookingsCollection.Find(ctx, filter)

	if err != nil {
		return payments, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return payments, err
	}

	for _, result := range results {
		cursor.Decode(&result)
		payments = append(payments, result.Payment)
	}

	return payments, err

}

func GetPaidPayment(username string) (payments []Payment, err error) {
	filter := bson.M{"trainer": username, "payment.status": "paid"}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var results []Booking

	cursor, err := bookingsCollection.Find(ctx, filter)

	if err != nil {
		return payments, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return payments, err
	}

	for _, result := range results {
		cursor.Decode(&result)
		payments = append(payments, result.Payment)
	}

	return payments, err

}
