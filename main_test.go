package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"trainder-api/controllers"
	"trainder-api/responses"
	"trainder-api/models"

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
	SubAddress  string `json:"subAddress" binding:"required"`
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestRegisterHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/register", controllers.Register())
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
		SubAddress:  "subaddress01",
	}
	jsonValue, _ := json.Marshal(registerInput)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLoginHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/login", controllers.Login())
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

	r.POST("/login", controllers.Login())
	r.GET("/protected/user", controllers.CurrentUser())
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
