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

type Message struct {
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	Content   string    `bson:"content" json:"content"`
	Sender    string    `bson:"sender" json:"sender"`
}

type Chat struct {
	RoomID   string    `bson:"roomID" json:"roomID"`
	Trainer  string    `bson:"trainer" json:"trainer"`
	Trainee  string    `bson:"trainee" json:"trainee"`
	Messages []Message `bson:"messages" json:"messages"`
}

type AllChat struct {
	Audience string  `bson:"audience" json:"audience"`
	Message  Message `bson:"message" json:"message"`
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
		"messages": []Message{},
	}
	_, err := chatCollection.InsertOne(ctx, chat)
	if err != nil {
		return fmt.Errorf("failed to init chat: %v", err)
	}

	return nil

}

func AddMessage(roomID string, content string, sender string) error {

	s := strings.Split(roomID, "_")
	trainer := s[1]
	trainee := s[3]

	chatExists, err := chatExists(trainer, trainee)
	if err != nil {
		return fmt.Errorf("failed at chatExists: %v", err)
	}
	if !chatExists {
		err = InitChat(trainer, trainee)
		if err != nil {
			return fmt.Errorf("failed at AddMessage: %v", err)
		}
	}

	message := Message{
		CreatedAt: time.Now(),
		Content:   content,
		Sender:    sender,
	}
	filter := bson.M{
		"roomID": roomID,
	}
	update := bson.M{"$push": bson.M{"messages": message}}
	_, err = chatCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil

}

func GetAllChatLatestMessage(username string) ([]AllChat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var filter bson.M
	trainerFlag := false
	if IsTrainer(username) {
		trainerFlag = true
		filter = bson.M{
			"trainer": username,
		}
	} else {
		filter = bson.M{
			"trainee": username,
		}
	}

	cursor, err := chatCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var allChats []AllChat
	for cursor.Next(ctx) {
		var chat Chat
		err := cursor.Decode(&chat)
		if err != nil {
			return nil, err
		}
		message := chat.Messages[len(chat.Messages)-1]
		var result AllChat
		if trainerFlag {
			result = AllChat{
				Audience: chat.Trainee,
				Message:  message,
			}
		} else {
			result = AllChat{
				Audience: chat.Trainer,
				Message:  message,
			}
		}

		allChats = append(allChats, result)
	}

	return allChats, nil
}

func GetPastChat(username string, audience string) ([]Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var filter bson.M

	var messages []Message
	if IsTrainer(username) {
		filter = bson.M{
			"trainer": username,
			"trainee": audience,
		}
	} else {
		filter = bson.M{
			"trainer": audience,
			"trainee": username,
		}
	}
	cursor, err := chatCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var chat Chat
		err := cursor.Decode(&chat)
		if err != nil {
			return nil, err
		}
		messages = chat.Messages
	}
	return messages, nil
}
