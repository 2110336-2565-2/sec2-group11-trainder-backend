// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "login with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "put login input and pass to  gin.Context",
                        "name": "json_in_ginContext",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.LoginResponse"
                        }
                    }
                }
            }
        },
        "/protected/add-review": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add review on trainer to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trainer"
                ],
                "summary": "Add trainer review",
                "parameters": [
                    {
                        "description": "Parameters for trainer review",
                        "name": "ReviewRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ReviewDetails"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.AddReviewResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.AddReviewResponse"
                        }
                    }
                }
            }
        },
        "/protected/bookings": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a list of upcoming bookings for the trainer who is currently logged in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get bookings for the logged in trainer",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.GetBookingsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.GetBookingsResponse"
                        }
                    }
                }
            }
        },
        "/protected/create-booking": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates a new booking with the specified trainer, trainee, date, start time, and end time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Create a new booking",
                "parameters": [
                    {
                        "description": "put booking details and pass to gin.Context",
                        "name": "json_in_ginContext",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.BookingForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "booking created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/protected/delete-booking": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a booking with the specified bookingId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Delete a booking",
                "parameters": [
                    {
                        "description": "put DeleteBookingForm details and pass to gin.Context",
                        "name": "json_in_ginContext",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.DeleteBookingForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully delete booking",
                        "schema": {
                            "$ref": "#/definitions/responses.DeleteBookingResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request, missing filed of objectId or cannot find bookingObjectId",
                        "schema": {
                            "$ref": "#/definitions/responses.DeleteBookingResponse"
                        }
                    }
                }
            }
        },
        "/protected/filter-trainer": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "FilterTrainer base on filter input",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trainer"
                ],
                "summary": "FilterTrainer base on filter input",
                "parameters": [
                    {
                        "description": "Parameters for filtering trainers",
                        "name": "FilterTrainer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.FilterTrainerForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.FilterTrainerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.FilterTrainerResponse"
                        }
                    }
                }
            }
        },
        "/protected/profile": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "getProfile of the current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "getProfile of the current user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized, the user is not logged in",
                        "schema": {
                            "$ref": "#/definitions/responses.UserProfileResponse"
                        }
                    }
                }
            }
        },
        "/protected/reviews": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get reviews of specific trainer username from database sort by recent date then rating desc, limit number of output by limit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trainer"
                ],
                "summary": "Get reviews of specific trainer",
                "parameters": [
                    {
                        "description": "Parameters for querying trainer reviews",
                        "name": "GetReviewsInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.GetReviewsForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.TrainerReviewsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.TrainerReviewsResponse"
                        }
                    }
                }
            }
        },
        "/protected/trainer": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves the trainer profile information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trainer"
                ],
                "summary": "Retrieve trainer profile",
                "parameters": [
                    {
                        "description": "Put username input for retrieving the trainer profile",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.GetTrainerForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved the trainer profile",
                        "schema": {
                            "$ref": "#/definitions/responses.TrainerProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Failed to retrieve the trainer profile",
                        "schema": {
                            "$ref": "#/definitions/responses.TrainerProfileResponse"
                        }
                    }
                }
            }
        },
        "/protected/trainer-profile": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves the trainer profile information of the current user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trainer"
                ],
                "summary": "Retrieve trainer profile of current user",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved the trainer profile",
                        "schema": {
                            "$ref": "#/definitions/responses.TrainerProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Failed to retrieve the trainer profile",
                        "schema": {
                            "$ref": "#/definitions/responses.TrainerProfileResponse"
                        }
                    }
                }
            }
        },
        "/protected/update-booking": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update a booking of sepecified bookingId with the specified update input consist of status(pending/confirm/complete) and paymentStatus(pending/paid)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Update a booking",
                "parameters": [
                    {
                        "description": "put updateBookingForm details and pass to gin.Context",
                        "name": "json_in_ginContext",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateBookingForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully update booking",
                        "schema": {
                            "$ref": "#/definitions/responses.UpdateBookingResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request, missing filed of objectId or cannot find bookingObjectId",
                        "schema": {
                            "$ref": "#/definitions/responses.UpdateBookingResponse"
                        }
                    }
                }
            }
        },
        "/protected/update-profile": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "updateProfile of the current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "updateProfile of the current user",
                "parameters": [
                    {
                        "description": "put profile input json and pass to  gin.Context",
                        "name": "ProfileToUpdate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ProfileDetails"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ProfileResponse"
                        }
                    }
                }
            }
        },
        "/protected/update-trainer": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trainer"
                ],
                "summary": "Update the trainer's profile information.",
                "parameters": [
                    {
                        "description": "Trainer's information to update",
                        "name": "profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateTrainerDetails"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully update the trainer's profile",
                        "schema": {
                            "$ref": "#/definitions/responses.ProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request, either invalid input or user is not a trainer",
                        "schema": {
                            "$ref": "#/definitions/responses.ProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized, the user is not logged in",
                        "schema": {
                            "$ref": "#/definitions/responses.ProfileResponse"
                        }
                    }
                }
            }
        },
        "/protected/user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    },
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get the current user's username.  After getting token replied from logging in, put token in ginContext's token field",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "get the current user's username",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.CurrentUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.CurrentUserResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register with username,password,UserType [\"trainer\",\"trainee\"],Firstname,Lastname,Birthdate (\"yyyy-mm-dd\"),CitizenId (len == 13),Gender [\"Male\",\"Female\",\"Other\"],PhoneNumber (len ==10),Address,SubAddress",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Register user",
                "operationId": "register-user",
                "parameters": [
                    {
                        "description": "put register input and pass to gin.Context",
                        "name": "json_in_ginContext",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.RegisterResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.BookingForm": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "endTime": {
                    "type": "string"
                },
                "startTime": {
                    "type": "string"
                },
                "trainer": {
                    "type": "string"
                }
            }
        },
        "controllers.DeleteBookingForm": {
            "type": "object",
            "required": [
                "bookingId"
            ],
            "properties": {
                "bookingId": {
                    "type": "string"
                }
            }
        },
        "controllers.FilterTrainerForm": {
            "type": "object",
            "required": [
                "limit"
            ],
            "properties": {
                "feeMax": {
                    "type": "number"
                },
                "feeMin": {
                    "type": "number"
                },
                "limit": {
                    "type": "integer"
                },
                "specialty": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "controllers.GetReviewsForm": {
            "type": "object",
            "required": [
                "limit",
                "trainerUsername"
            ],
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "trainerUsername": {
                    "type": "string"
                }
            }
        },
        "controllers.GetTrainerForm": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.ProfileDetails": {
            "type": "object",
            "required": [
                "address",
                "birthdate",
                "citizenId",
                "firstname",
                "gender",
                "lastname",
                "phoneNumber"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "avatarUrl": {
                    "type": "string"
                },
                "birthdate": {
                    "type": "string"
                },
                "citizenId": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                }
            }
        },
        "controllers.RegisterInput": {
            "type": "object",
            "required": [
                "address",
                "birthdate",
                "citizenId",
                "firstname",
                "gender",
                "lastname",
                "lat",
                "lng",
                "password",
                "phoneNumber",
                "username",
                "usertype"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "avatarUrl": {
                    "type": "string"
                },
                "birthdate": {
                    "type": "string"
                },
                "citizenId": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "lat": {
                    "type": "number"
                },
                "lng": {
                    "type": "number"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "usertype": {
                    "type": "string"
                }
            }
        },
        "controllers.ReviewDetails": {
            "type": "object",
            "required": [
                "comment",
                "rating",
                "trainerUsername"
            ],
            "properties": {
                "comment": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "trainerUsername": {
                    "type": "string"
                }
            }
        },
        "controllers.UpdateBookingForm": {
            "type": "object",
            "required": [
                "bookingId"
            ],
            "properties": {
                "bookingId": {
                    "type": "string"
                },
                "paymentStatus": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "controllers.UpdateTrainerDetails": {
            "type": "object",
            "properties": {
                "certificateUrl": {
                    "type": "string"
                },
                "fee": {
                    "type": "number"
                },
                "rating": {
                    "type": "number"
                },
                "specialty": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "traineeCount": {
                    "type": "integer"
                }
            }
        },
        "models.Booking": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "endDateTime": {
                    "type": "string"
                },
                "payment": {
                    "type": "object",
                    "properties": {
                        "status": {
                            "type": "string"
                        },
                        "totalCost": {
                            "type": "number"
                        }
                    }
                },
                "startDateTime": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "trainee": {
                    "type": "string"
                },
                "trainer": {
                    "type": "string"
                }
            }
        },
        "models.FilteredTrainerInfo": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "avatarUrl": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "trainerInfo": {
                    "$ref": "#/definitions/models.TrainerInfo"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Review": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.TrainerInfo": {
            "type": "object",
            "properties": {
                "certificateURL": {
                    "type": "string"
                },
                "fee": {
                    "type": "number"
                },
                "rating": {
                    "type": "number"
                },
                "specialty": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "traineeCount": {
                    "type": "integer"
                }
            }
        },
        "models.UserProfile": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "avatarUrl": {
                    "type": "string"
                },
                "birthdate": {
                    "type": "string"
                },
                "citizenId": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "lat": {
                    "type": "number"
                },
                "lng": {
                    "type": "number"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "usertype": {
                    "type": "string"
                }
            }
        },
        "responses.AddReviewResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.CurrentUserResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "responses.DeleteBookingResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.FilterTrainerResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "trainers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.FilteredTrainerInfo"
                    }
                }
            }
        },
        "responses.GetBookingsResponse": {
            "type": "object",
            "properties": {
                "bookings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Booking"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.LoginResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "responses.ProfileResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.RegisterResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.TrainerProfileResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "trainerInfo": {
                    "$ref": "#/definitions/models.TrainerInfo"
                },
                "user": {
                    "$ref": "#/definitions/models.UserProfile"
                }
            }
        },
        "responses.TrainerReviewsResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "reviews": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Review"
                    }
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.UpdateBookingResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.UserProfileResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/models.UserProfile"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Trainder API",
	Description:      "API for Trainder",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
