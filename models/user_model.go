package models

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindProfile(username string) (resultMap map[string]interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result bson.M
	err = userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&result)
	if err != nil {
		return nil, err
	}

	delete(result, "_id")
	delete(result, "createdAt")
	delete(result, "updatedAt")
	delete(result, "hashedPassword")
	resultMap = result

	return resultMap, nil

}

func UpdateUserProfile(username string, firstName string, lastName string, birthDate string, citizenID string, gender string, phoneNumber string, address string, subAddress string) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// err = ProfileConditionCheck(firstName, lastName, birthDate, citizenID, gender, phoneNumber)

	layout := "2006-01-02"
	date, error := time.Parse(layout, birthDate)

	if error != nil {
		return
	}

	result, err = userCollection.UpdateOne(
		ctx,
		bson.M{"username": username},
		bson.M{"$set": bson.M{
			"firstname":   firstName,
			"lastname":    lastName,
			"birthdate":   date,
			"citizenId":   citizenID,
			"gender":      gender,
			"phoneNumber": phoneNumber,
			"address":     address,
			"subAddress":  subAddress,
			"updatedAt":   time.Now(),
		}},
	)
	return
}

func ProfileConditionCheck(firstName string, lastName string, birthDate string, citizenID string, gender string, phoneNumber string) error {
	//---------------check firstName not contain strange character
	//---------------check date format

	const dateFormat = `^\d{4}-\d{2}-\d{2}$`
	match, err := regexp.MatchString(dateFormat, birthDate)
	if err != nil {
		fmt.Println("xxxx")
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
	// ---------------check phoneNumber
	if len(phoneNumber) != 10 {
		return errors.New("phoneNumber is not valid, must have length 10")
	}

	return nil
}

// TODO: enable id checking for production" uncomment below
// func assertThaiID(thaiID string) error {
// 	re := regexp.MustCompile(`(\d{12})(\d)`)
// 	matches := re.FindStringSubmatch(thaiID)
// 	if len(matches) == 0 {
// 		return errors.New("bad input from user, invalid thaiID, length != 13")

// 	}

// 	digits := matches[1]
// 	sum := 0
// 	for i, digit := range digits {
// 		d, _ := strconv.Atoi(string(digit))
// 		sum += (13 - i) * d
// 	}
// 	lastDigit := (11 - sum%11) % 10
// 	inputLastDigit, _ := strconv.Atoi(matches[2])
// 	if lastDigit != inputLastDigit {
// 		return errors.New("bad input from user, invalid thaiID")

// 	}

// 	return nil
// }
