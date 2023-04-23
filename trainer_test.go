package main

import (
	"log"
	"math"
	"os"
	"testing"
	"trainder-api/models"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

type UserInfo struct {
	Username    string
	Password    string
	UserType    string
	FirstName   string
	LastName    string
	BirthDate   string
	CitizenId   string
	Gender      string
	PhoneNumber string
	Address     string
	Lat         float64
	Lng         float64
	AvatarUrl   string
}

func addTestUser() {
	testUserInfo := UserInfo{
		Username:    "__test",
		Password:    "123456789",
		UserType:    "Trainer",
		FirstName:   "",
		LastName:    "Trainder",
		BirthDate:   "2000-01-01",
		Gender:      "Other",
		PhoneNumber: "0000000000",
		Address:     "-",
		AvatarUrl:   "",
		Lat:         0,
		Lng:         0,
	}

	_, err := models.CreateUser(testUserInfo.Username, testUserInfo.Password,
		testUserInfo.UserType, testUserInfo.FirstName, testUserInfo.LastName,
		testUserInfo.BirthDate, testUserInfo.CitizenId, testUserInfo.Gender,
		testUserInfo.PhoneNumber, testUserInfo.Address, testUserInfo.AvatarUrl,
		testUserInfo.Lat, testUserInfo.Lng)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	var specialties []string
	specialties = append(specialties, "Weight Loss")
	specialties = append(specialties, "Rehabilitation")

	_, err = models.UpdateTrainerProfile(testUserInfo.Username, specialties, 0, 100, 0, "")
	if err != nil {
		log.Fatal(err.Error())
	}

}

func deleteTestUser() {
	err := models.DeleteUser("__test")
	if err != nil {
		panic(err.Error())
	}
}

func TestMain(m *testing.M) {
	addTestUser()
	code := m.Run()

	deleteTestUser()

	os.Exit(code)

}
func TestFilterTrainerWithLimit(t *testing.T) {
	var specialties []string

	result, err := models.FindFilteredTrainer(specialties, 1, 0, math.MaxInt)

	assert.Equal(t, nil, err)

	assert.Equal(t, 1, len(result))

}

func TestFilterTrainerWithNoFeeLimit(t *testing.T) {
	var specialties []string

	result, err := models.FindFilteredTrainer(specialties, 1, 0, 0)

	assert.Equal(t, nil, err)

	assert.Equal(t, 1, len(result))

}

func TestFilterTrainerWithOneSpecialty(t *testing.T) {
	var specialties []string
	specialties = append(specialties, "Weight Loss")

	result, err := models.FindFilteredTrainer(specialties, 1, 0, 0)

	assert.Equal(t, nil, err)

	assert.LessOrEqual(t, 1, len(result))

	if len(result) > 0 {
		done := false
		for _, spec := range result[0].TrainerInfo.Specialty {
			if slices.Contains(specialties, spec) {
				done = true
				break
			}
		}
		assert.True(t, done)
	}
}

func TestFilterTrainerWithManySpecialty(t *testing.T) {
	var specialties []string
	specialties = append(specialties, "Weight Loss")
	specialties = append(specialties, "Rehabilitation")

	result, err := models.FindFilteredTrainer(specialties, 1, 0, 0)

	assert.Equal(t, nil, err)

	assert.LessOrEqual(t, 1, len(result))

	if len(result) > 0 {
		done := false
		for _, spec := range result[0].TrainerInfo.Specialty {
			if slices.Contains(specialties, spec) {
				done = true
				break
			}
		}
		assert.True(t, done)
	}
}

func TestFilterTrainerWithInvalidSpecialty(t *testing.T) {
	var specialties []string
	specialties = append(specialties, "Bla Bla")

	result, err := models.FindFilteredTrainer(specialties, 1, 0, 0)

	assert.Equal(t, nil, err)

	assert.Equal(t, len(result), 0)

}
