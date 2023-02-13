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
                            "$ref": "#/definitions/controllers.FilterTrainerInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
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
                            "$ref": "#/definitions/controllers.GetTrainerInput"
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
                            "$ref": "#/definitions/controllers.ProfileInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
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
                            "$ref": "#/definitions/controllers.UpdateTrainerInput"
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
                        "description": "put register input and pass to  gin.Context",
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
        "controllers.FilterTrainerInput": {
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
        "controllers.GetTrainerInput": {
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
        "controllers.ProfileInput": {
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
        "controllers.UpdateTrainerInput": {
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
        "models.TrainerInfo": {
            "type": "object",
            "properties": {
                "certificateUrl": {
                    "type": "string"
                },
                "fee": {
                    "type": "integer"
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
