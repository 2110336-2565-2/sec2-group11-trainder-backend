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

type User struct {
	Username       string `bson:"username"`
	HashedPassword string `bson:"hashedPassword"`
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

func CreateUser(username string, password string) (user User, err error) {
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

	user = User{
		Username:       username,
		HashedPassword: string(hashedPassword),
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
