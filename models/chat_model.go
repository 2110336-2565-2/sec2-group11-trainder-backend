package models

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trainder-api/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var chatCollection *mongo.Collection = configs.GetCollection(configs.DB, "chats")

type Messege struct {
	CreatedAt time.Time `bson:"createdAt"`
	Content   string    `bson:"content"`
	Sender    string    `bson:"sender"`
}

type Chat struct {
	RoomID   string    `bson:"roomID" json:"roomID"`
	Trainer  string    `bson:"trainer" json:"trainer"`
	Trainee  string    `bson:"trainee" json:"trainee"`
	Messeges []Messege `bson:"messeges"`
}

// func FindChat(trainer string,trainee string) (chat Chat, err error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	filter := bson.M{
// 		"trainer": trainer,
// 		"trainee": trainee,
// 	}
// 	err = userCollection.FindOne(ctx, filter).Decode(&chat)
// 	return chat, err
// }

func chatExists(trainer string, trainee string) (bool, error) {
	filter := bson.M{
		"trainer": trainer,
		"trainee": trainee,
	}
	count, err := chatCollection.CountDocuments(context.Background(), filter, nil)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func InitChat(trainer string, trainee string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rid := fmt.Sprintf("trainer_%s_trainee_%s", trainer, trainee)
	chat := bson.M{
		"roomID":   rid,
		"trainer":  trainer,
		"trainee":  trainee,
		"messeges": []Messege{},
	}
	_, err := chatCollection.InsertOne(ctx, chat)
	if err != nil {
		return fmt.Errorf("failed to init chat: %v", err)
	}

	return nil

}

// func ddReview(trainerUsername string, username string, rating int, comment string) error {
// 	isExist, err := userExists(trainerUsername)
// 	if err != nil {
// 		return err
// 	}
// 	if !isExist {
// 		err = &UserNotExist{}
// 		return err
// 	}
// 	review := Review{
// 		Username:  username,
// 		Rating:    rating,
// 		Comment:   comment,
// 		CreatedAt: time.Now(),
// 	}
// 	filter := bson.M{"username": trainerUsername}
// 	update := bson.M{"$push": bson.M{"reviews": review}}
// 	_, err = userCollection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	err = updateRatingByUsername(trainerUsername)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

func AddMessege(roomID string, content string, sender string) error {

	s := strings.Split(roomID, "_")
	trainer := s[1]
	trainee := s[3]

	fmt.Println("AddMessege", trainer, trainee)
	// fmt.Println(ListCollectionNames())
	chatexists, err := chatExists(trainer, trainee)
	if err != nil {
		return fmt.Errorf("failed at chatExists: %v", err)
	}
	if !chatexists {
		err = InitChat(trainer, trainee)
		if err != nil {
			return fmt.Errorf("failed at AddMessege: %v", err)
		}
	}

	messege := Messege{
		CreatedAt: time.Now(),
		Content:   content,
		Sender:    sender,
	}
	// rid := fmt.Sprintf("trainer_%s_trainee_%s", trainer, trainee)
	filter := bson.M{
		"roomID": roomID,
	}
	update := bson.M{"$push": bson.M{"messeges": messege}}
	_, err = chatCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil

}
