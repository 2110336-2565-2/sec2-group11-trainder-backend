package controllers

import (
	"net/http"
	"trainder-api/models"
	"trainder-api/responses"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	UserType    string `json:"usertype" binding:"required"`
	Firstname   string `json:"firstname" binding:"required"`
	Lastname    string `json:"lastname" binding:"required"`
	Birthdate   string `jason:"birthdate" binding:"required"`
	CitizenId   string `jason:"citizenid" binding:"required"`
	Gender      string `jason:"gender" binding:"required"`
	PhoneNumber string `jason:"phonenumber" binding:"required"`
	Address     string `json:"addresss" binding:"required"`
	SubAddress  string `json:"subaddresss" binding:"required"`
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {

		var input RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.RegisterResponse{Status: http.StatusBadRequest, Message: "input missing"})
			return
		}

		_, err := models.CreateUser(input.Username, input.Password, input.UserType, input.Firstname, input.Lastname, input.Birthdate, input.CitizenId, input.Gender, input.PhoneNumber, input.Address, input.SubAddress)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.RegisterResponse{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, responses.RegisterResponse{Status: http.StatusCreated, Message: "registration success!"})
	}
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error"})
			return
		}

		user, err := models.FindUser(input.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "user not found"})
			return
		}

		token, err := user.LoginCheck(input.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "password is incorrect"})
			return
		}
		c.JSON(http.StatusOK, responses.LoginResponse{Status: http.StatusOK, Message: "logged in", Token: token})
	}

}
