package models

import (
	"context"
	"time"
	"trainder-api/configs"
	"trainder-api/utils/tokens"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

type TrainerInfo struct {
	Specialty      []string `bson:"specialty"      json:"specialty"`
	Fee            int      `bson:"fee,omitempty"            json:"fee"`
	CertificateURL string   `bson:"certificateUrl,omitempty" json:"certificateUrl"`
	Rating         float64  `bson:"rating,omitempty"         json:"rating"`
	TraineeCount   int32    `bson:"traineeCount,omitempty"   json:"traineeCount"`
}
type User struct {
	Username       string      `bson:"username"`
	HashedPassword string      `bson:"hashedPassword"`
	UserType       string      `bson:"usertype"`
	FirstName      string      `bson:"firstname"`
	LastName       string      `bson:"lastname"`
	BirthDate      time.Time   `bson:"birthdate"`
	CitizenId      string      `bson:"citizenId"`
	Gender         string      `bson:"gender"`
	PhoneNumber    string      `bson:"phoneNumber"`
	Address        string      `bson:"address"`
	CreatedAt      time.Time   `bson:"createdAt"`
	UpdatedAt      time.Time   `bson:"updatedAt"`
	AvatarUrl      string      `bson:"avatarUrl"`
	Lat            float64     `bson:"lat"`
	Lng            float64     `bson:"lng"`
	TrainerInfo    TrainerInfo `bson:"trainerInfo,omitempty"`
}

func (tr TrainerInfo) Init() TrainerInfo {
	tr.Fee = 200
	tr.CertificateURL = "certificateURLString"
	tr.Rating = 3
	tr.TraineeCount = 0
	tr.Specialty = []string{}
	return tr
}

func FindUser(username string) (user User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "username", Value: username}}
	err = userCollection.FindOne(ctx, filter).Decode(&user)
	return user, err
}

type PasswordConditionCheckFailed struct {
	s string
}

func (e *PasswordConditionCheckFailed) Error() string {
	return e.s
}

func passwordConditionCheck(password string) error {
	if len(password) < 8 {
		return &PasswordConditionCheckFailed{"Password should be more than 8 characters"}
	}
	return nil
}

type UserAlreadyExist struct{}

func (e *UserAlreadyExist) Error() string {
	return "error: user already existed"
}

func CreateUser(username string, password string, userType string, firstName string, lastName string, birthDate string, citizenID string, gender string, phoneNumber string, address string, avatarUrl string, lat float64, lng float64) (user User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user, err = FindUser(username)
	if err == nil {
		err = &UserAlreadyExist{}
		return user, err
	}
	err = passwordConditionCheck(password)
	if err != nil {
		return user, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	layout := "2006-01-02"
	date, error := time.Parse(layout, birthDate)

	if error != nil {
		return user, err
	}
	var initTrainer TrainerInfo
	if userType == "Trainer" {
		initTrainer = new(TrainerInfo).Init()
		user = User{
			Username:       username,
			HashedPassword: string(hashedPassword),
			UserType:       userType,
			FirstName:      firstName,
			LastName:       lastName,
			BirthDate:      date,
			CitizenId:      citizenID,
			Gender:         gender,
			PhoneNumber:    phoneNumber,
			Address:        address,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			AvatarUrl:      avatarUrl,
			Lat:            lat,
			Lng:            lng,
			TrainerInfo:    initTrainer,
		}

	} else {
		user = User{
			Username:       username,
			HashedPassword: string(hashedPassword),
			UserType:       userType,
			FirstName:      firstName,
			LastName:       lastName,
			BirthDate:      date,
			CitizenId:      citizenID,
			Gender:         gender,
			PhoneNumber:    phoneNumber,
			Address:        address,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Lat:            lat,
			Lng:            lng,
			AvatarUrl:      avatarUrl,
		}
	}

	_, err = userCollection.InsertOne(ctx, user)
	return user, err
}

func (user *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
}

func (user *User) LoginCheck(password string) (token string, err error) {
	err = user.VerifyPassword(password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err = tokens.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, err
}

func DeleteUser(username string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "username", Value: username}}
	_, err = userCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
