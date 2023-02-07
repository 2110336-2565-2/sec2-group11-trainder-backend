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
                "subAddress",
                "username",
                "usertype"
            ],
            "properties": {
                "address": {
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
                "subAddress": {
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
