package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"trainder-api/models"
	"trainder-api/responses"
	"trainder-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	UserType    string `json:"usertype" binding:"required"`
	Firstname   string `json:"firstname" binding:"required"`
	Lastname    string `json:"lastname" binding:"required"`
	Birthdate   string `json:"birthdate" binding:"required"`
	CitizenId   string `json:"citizenId" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type UpdateUser struct {
	UserType    string `json:"usertype" binding:"required"`
	Firstname   string `json:"firstname" binding:"required"`
	Lastname    string `json:"lastname" binding:"required"`
	Birthdate   string `json:"birthdate" binding:"required"`
	CitizenId   string `json:"citizenId" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type UpdateTrainer struct {
	Specialty      []string `json:"specialty"`
	Rating         float64  `json:"rating"`
	Fee            float64  `json:"fee"`
	TraineeCount   int32    `json:"traineeCount"`
	CertificateUrl string   `json:"certificateUrl"`
}

type FilterTrainer struct {
	Specialty []string `json:"specialty"`
	Limit     int      `json:"limit" binding:"required"`
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestRegisterHandler(t *testing.T) {
	r := SetUpRouter()
	routes.AuthRoute(r)
	routes.ProtectedRoute(r)
	_ = models.DeleteUser("test01")
	registerInput := RegisterInput{
		Username:    "test01",
		Password:    "password01",
		UserType:    "Trainer",
		Firstname:   "firstname",
		Lastname:    "lastname",
		Birthdate:   "2022-01-12",
		CitizenId:   "9261991922738",
		Gender:      "Male",
		PhoneNumber: "0881234567",
		Address:     "address01",
	}
	jsonValue, _ := json.Marshal(registerInput)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLoginHandler(t *testing.T) {
	r := SetUpRouter()
	routes.AuthRoute(r)
	routes.ProtectedRoute(r)
	user := User{
		Username: "test01",
		Password: "password01",
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestUserHandler(t *testing.T) {
	r := SetUpRouter()
	routes.AuthRoute(r)
	routes.ProtectedRoute(r)
	user := User{
		Username: "test01",
		Password: "password01",
	}
	jsonValue, _ := json.Marshal(user)
	loginReq, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, loginReq)

	var response responses.LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response: ", err)
	}
	jwt := response.Token
	userReq, _ := http.NewRequest("GET", "/protected/user", nil)
	userReq.Header.Set("Authorization", "Bearer "+jwt)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, userReq)

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestUpdateHanlder(t *testing.T) {
	r := SetUpRouter()
	routes.AuthRoute(r)
	routes.ProtectedRoute(r)
	user := User{
		Username: "test01",
		Password: "password01",
	}
	update := UpdateUser{
		UserType:    "Trainer",
		Firstname:   "firstname",
		Lastname:    "lastname",
		Birthdate:   "2022-01-12",
		CitizenId:   "9261991922738",
		Gender:      "Male",
		PhoneNumber: "0881234567",
		Address:     "address01",
	}
	jsonValue, _ := json.Marshal(user)
	loginReq, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, loginReq)

	var response responses.LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response: ", err)
	}
	jwt := response.Token
	jsonValue, _ = json.Marshal(update)
	userReq, _ := http.NewRequest("POST", "/protected/update-profile", bytes.NewBuffer(jsonValue))
	userReq.Header.Set("Authorization", "Bearer "+jwt)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, userReq)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTrainerHandler(t *testing.T) {
	r := SetUpRouter()
	routes.AuthRoute(r)
	routes.ProtectedRoute(r)
	user := User{
		Username: "test01",
		Password: "password01",
	}
	update := UpdateTrainer{
		Specialty:      []string{"yoga", "pilates"},
		Rating:         2.0,
		Fee:            200.0,
		TraineeCount:   2,
		CertificateUrl: "url",
	}
	jsonValue, _ := json.Marshal(user)
	loginReq, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, loginReq)

	var response responses.LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response: ", err)
	}
	jwt := response.Token
	jsonValue, _ = json.Marshal(update)
	userReq, _ := http.NewRequest("POST", "/protected/update-trainer", bytes.NewBuffer(jsonValue))
	userReq.Header.Set("Authorization", "Bearer "+jwt)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, userReq)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFilter(t *testing.T) {
	r := SetUpRouter()
	routes.AuthRoute(r)
	routes.ProtectedRoute(r)
	user := User{
		Username: "test01",
		Password: "password01",
	}
	update := FilterTrainer{
		Specialty: []string{"yoga"},
		Limit:     1,
	}
	jsonValue, _ := json.Marshal(user)
	loginReq, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, loginReq)

	var response responses.LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("Failed to parse response: ", err)
	}
	jwt := response.Token
	jsonValue, _ = json.Marshal(update)
	userReq, _ := http.NewRequest("POST", "/protected/filter-trainer", bytes.NewBuffer(jsonValue))
	userReq.Header.Set("Authorization", "Bearer "+jwt)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, userReq)
	assert.Equal(t, http.StatusOK, w.Code)
}
