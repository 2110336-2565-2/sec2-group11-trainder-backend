package models

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserProfile struct {
	Username    string  `json:"username"`
	UserType    string  `json:"usertype"`
	FirstName   string  `json:"firstname"`
	LastName    string  `json:"lastname"`
	BirthDate   string  `json:"birthdate"`
	CitizenId   string  `json:"citizenId"`
	Gender      string  `json:"gender"`
	PhoneNumber string  `json:"phoneNumber"`
	Address     string  `json:"address"`
	AvatarUrl   string  `json:"avatarUrl"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
}

func FindProfile(username, userType string) (result UserProfile, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "username", Value: username}}
	if userType != "" {
		filter = append(filter, bson.E{Key: "usertype", Value: userType})
	}
	opts := options.FindOne().SetProjection(bson.D{
		{Key: "_id", Value: 0},
		{Key: "hashedPassword", Value: 0},
		{Key: "createdAt", Value: 0},
		{Key: "updatedAt", Value: 0}})
	var user User
	err = userCollection.FindOne(ctx, filter, opts).Decode(&user)
	if err != nil {
		return result, err
	}
	result = UserProfile{
		Username:    user.Username,
		UserType:    user.UserType,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		BirthDate:   user.BirthDate.Format("2006-01-02"),
		CitizenId:   user.CitizenId,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		AvatarUrl:   user.AvatarUrl,
	}
	return result, nil
}

func UpdateUserProfile(username string, firstName string, lastName string, birthDate string, citizenID string, gender string, phoneNumber string, address string, avatarUrl string) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// err = ProfileConditionCheck(firstName, lastName, birthDate, citizenID, gender, phoneNumber)

	layout := "2006-01-02"
	date, error := time.Parse(layout, birthDate)

	if error != nil {
		return
	}

	info := bson.M{
		"firstname":   firstName,
		"lastname":    lastName,
		"birthdate":   date,
		"citizenId":   citizenID,
		"gender":      gender,
		"phoneNumber": phoneNumber,
		"address":     address,
		"updatedAt":   time.Now(),
	}
	if avatarUrl != "" {
		info["avatarUrl"] = avatarUrl
	}

	result, err = userCollection.UpdateOne(
		ctx,
		bson.M{"username": username},
		bson.M{"$set": info},
	)
	return
}

func UpdateAvatarUrl(username string, imageID string) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{
		"updatedAt": time.Now(),
		"avatarUrl": imageID,
	}}
	result, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

func GetAvatarUrl(username string) (string, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// filter := bson.M{"username": username}
	profile, err := FindProfile(username, "")
	if err != nil {
		return "", fmt.Errorf("Error from FindProfile in GetAvatarUrl %v", err)
	}
	// fmt.Println("GetAvatarUrl", profile.AvatarUrl)

	return profile.AvatarUrl, nil
}

func ProfileConditionCheck(firstName string, lastName string, birthDate string, citizenID string, gender string, phoneNumber string) error {
	//---------------check firstName not contain strange character, number  (accepting)
	if alpha := isAlphaString(firstName); !alpha {
		return errors.New("firstName invalid")
	}

	if alpha := isAlphaString(lastName); !alpha {
		return errors.New("lastName invalid")
	}
	//---------------check date format

	const dateFormat = `^\d{4}-\d{2}-\d{2}$`
	match, err := regexp.MatchString(dateFormat, birthDate)
	if err != nil {
		// fmt.Println("xxxx")
		return err
	}
	if !match {
		return errors.New("date must be in the format 'yyyy-mm-dd'")
	}
	//---------------check date value
	layout := "2006-01-02"
	date, error := time.Parse(layout, birthDate)
	if error != nil {
		return error
	}
	now := time.Now()
	if date.After(now) {
		return errors.New("date is After current time")

	}
	// ---------------check citizenID
	// TODO: enable id checking for production"
	// err = assertThaiID(citizenID)
	// if err != nil {
	// 	return err
	// }

	// ---------------check gender
	validGenders := []string{"Male", "Female", "Other"}

	isValidGender := false
	for _, v := range validGenders {
		if v == gender {
			isValidGender = true
			break
		}
	}

	if !isValidGender {
		return errors.New(" gender is not valid, valid gender in ['Male', 'Female', 'Other'] ")
	}
	// check phoneNumber
	if len(phoneNumber) != 10 {
		return errors.New("phoneNumber is not valid, must have length 10")
	}

	return nil
}

func assertThaiID(thaiID string) error {
	re := regexp.MustCompile(`(\d{12})(\d)`)
	matches := re.FindStringSubmatch(thaiID)
	if len(matches) == 0 {
		return errors.New("bad input from user, invalid thaiID, length != 13")

	}

	digits := matches[1]
	sum := 0
	for i, digit := range digits {
		d, _ := strconv.Atoi(string(digit))
		sum += (13 - i) * d
	}
	lastDigit := (11 - sum%11) % 10
	inputLastDigit, _ := strconv.Atoi(matches[2])
	if lastDigit != inputLastDigit {
		return errors.New("bad input from user, invalid thaiID")

	}

	return nil
}

func isAlphaString(input string) bool {
	// Use a regular expression to match only alphabet characters
	match, _ := regexp.MatchString("^[a-zA-Z\u0E00-\u0E7F]+$", input)
	return match
}

// use for normal registration the user cannot be "Admin"
func UserTypeCheck(userType string) error {
	validUserType := []string{"Trainer", "Trainee"}

	isValidUserType := false
	for _, v := range validUserType {
		if v == userType {
			isValidUserType = true
			break
		}
	}

	if !isValidUserType {
		return errors.New(" userType is not valid, valid userType in ['Trainer', 'Trainee'] ")
	}
	return nil
}

type NameAndRole struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserType  string `json:"usertype"`
}

func GetNameAndRole(username string) (NameAndRole, error) {
	var result NameAndRole
	profile, err := FindProfile(username, "")
	if err != nil {
		return result, fmt.Errorf("Error from FindProfile in GetName %v", err)
	}
	result = NameAndRole{
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		UserType:  profile.UserType,
	}
	return result, nil

}
