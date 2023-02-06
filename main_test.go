package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"trainder-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestLoginHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/login", controllers.Login())
	user := User{
		Username: "test10",
		Password: "password0",
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
