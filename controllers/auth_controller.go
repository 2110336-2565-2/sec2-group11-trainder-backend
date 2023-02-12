package controllers

import (
	"fmt"
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
	Birthdate   string `json:"birthdate" binding:"required"`
	CitizenId   string `json:"citizenId" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address     string `json:"address" binding:"required"`
	SubAddress  string `json:"subAddress" binding:"required"`
	AvatarUrl   string `json:"avatarUrl"`
}

//	@Summary		Register user
//	@Description	Register with username,password,UserType ["trainer","trainee"],Firstname,Lastname,Birthdate ("yyyy-mm-dd"),CitizenId (len == 13),Gender ["Male","Female","Other"],PhoneNumber (len ==10),Address,SubAddress
//	@Tags			authentication
//	@ID				register-user
//	@Accept			json
//	@Produce		json
//
//	@Param			json_in_ginContext	body		RegisterInput	true	"put register input and pass to  gin.Context"
//
//	@Success		200					{object}	responses.RegisterResponse
//	@Router			/register [post]
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {

		var input RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, responses.RegisterResponse{Status: http.StatusBadRequest, Message: "input missing"})
			return
		}

		profileErr := models.ProfileConditionCheck(input.Firstname, input.Lastname, input.Birthdate, input.CitizenId, input.Gender, input.PhoneNumber)
		if profileErr != nil {
			c.JSON(http.StatusBadRequest, responses.RegisterResponse{
				Status:  http.StatusBadRequest,
				Message: profileErr.Error(),
			})
			return
		}
		userTypeErr := models.UserTypeCheck(input.UserType)
		if userTypeErr != nil {
			c.JSON(http.StatusBadRequest, responses.RegisterResponse{
				Status:  http.StatusBadRequest,
				Message: userTypeErr.Error(),
			})
			return
		}

		_, err := models.CreateUser(input.Username, input.Password, input.UserType, input.Firstname, input.Lastname, input.Birthdate, input.CitizenId, input.Gender, input.PhoneNumber, input.Address, input.SubAddress, input.AvatarUrl)

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

// Login godoc
//
//	@Summary		Login
//	@Description	login with username and password
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			json_in_ginContext	body		LoginInput	true	"put login input and pass to  gin.Context"
//	@Success		200					{object}	responses.LoginResponse
//	@Router			/login [post]
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
		c.JSON(http.StatusOK, responses.LoginResponse{Status: http.StatusOK, Message: "logged in", Token: token, Username: input.Username})
	}
}
