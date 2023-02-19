package inits

import (
	"context"
	"fmt"
	"log"
	"time"
	"trainder-api/configs"
	"trainder-api/models"

	"go.mongodb.org/mongo-driver/bson"
)

func InitializeDatabase() {
	// Check if `users` collection exists
	db := configs.DB.Database(configs.EnvMongoDBName())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collectionNames, err := db.ListCollectionNames(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	userCollectionExists := false
	for _, collectionName := range collectionNames {
		if collectionName == "users" {
			userCollectionExists = true
		}
	}
	if !userCollectionExists {
		fmt.Println("User collections does not exist, creating `users` collection with `root` user")
		CreateRootUser()
	}
}

type RootUserInfo struct {
	Username    string
	Password    string
	UserType    string
	Firstname   string
	Lastname    string
	Birthdate   string
	CitizenId   string
	Gender      string
	PhoneNumber string
	Address     string
	Lat         float64
	Lng         float64
	AvatarUrl   string
}

// This function will be run to populate the databases with an initial user
// The user will be `root` and that users password will be defined as env ""
func CreateRootUser() {
	rootUserInfo := RootUserInfo{
		Username:    "root",
		Password:    configs.EnvInitRootPassword(),
		UserType:    "Admin",
		Firstname:   "Admin",
		Lastname:    "Trainder",
		Birthdate:   time.Now().Format("2006-01-02"),
		CitizenId:   "0000000000000",
		Gender:      "Other",
		PhoneNumber: "0000000000",
		Address:     "-",
		Lat:         0,
		Lng:         0,
		AvatarUrl:   "",
	}

	profileErr := models.ProfileConditionCheck(rootUserInfo.Firstname, rootUserInfo.Lastname, rootUserInfo.Birthdate, rootUserInfo.CitizenId, rootUserInfo.Gender, rootUserInfo.PhoneNumber)
	if profileErr != nil {
		log.Fatal(profileErr.Error())
		return
	}

	_, err := models.CreateUser(rootUserInfo.Username, rootUserInfo.Password, rootUserInfo.UserType, rootUserInfo.Firstname, rootUserInfo.Lastname, rootUserInfo.Birthdate, rootUserInfo.CitizenId, rootUserInfo.Gender, rootUserInfo.PhoneNumber, rootUserInfo.Address, rootUserInfo.AvatarUrl, rootUserInfo.Lat, rootUserInfo.Lng)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
