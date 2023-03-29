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

type AllChat struct {
	Audience string  `json:"audience"`
	Messege  Messege `bson:"messege"`
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

func GetAllChatLatestMessege(username string) ([]AllChat, error) {
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
	// fmt.Println(cursor)
	var allChats []AllChat
	for cursor.Next(ctx) {
		var chat Chat
		err := cursor.Decode(&chat)
		if err != nil {
			return nil, err
		}
		// fmt.Println("chat.Messeges", chat.Messeges, chat.Messeges[len(chat.Messeges)-1])
		messege := chat.Messeges[len(chat.Messeges)-1]
		// fmt.Println(messege.Content, messege.CreatedAt)
		var result AllChat
		if trainerFlag {
			result = AllChat{
				Audience: chat.Trainee,
				Messege:  messege,
			}
		} else {
			result = AllChat{
				Audience: chat.Trainer,
				Messege:  messege,
			}
		}

		allChats = append(allChats, result)
		fmt.Println("allChats", allChats)
	}

	return allChats, nil
}

func GetPastChat(username string, audience string) (Chat, error) {
	var chat Chat
	return chat, nil
}
