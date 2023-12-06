// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/credentials/login": {
            "post": {
                "description": "Handles log in for new users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Credentials"
                ],
                "summary": "Log in handler",
                "parameters": [
                    {
                        "description": "Login request body",
                        "name": "logInRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token response consisting of access and refresh token",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Data"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.TokenResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/credentials/refresh": {
            "post": {
                "description": "Handles requesting a new set of access and refresh token from the given previous refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Credentials"
                ],
                "summary": "Refresh handler",
                "parameters": [
                    {
                        "description": "Refresh token request body",
                        "name": "refreshRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token response consisting of access and refresh token",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Data"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.TokenResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/credentials/signup": {
            "post": {
                "description": "Handles sign up for new users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Credentials"
                ],
                "summary": "Sign up handler",
                "parameters": [
                    {
                        "description": "Signup request body",
                        "name": "signUpRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token response consisting of access and refresh token",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Data"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.TokenResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/ptd/profiles/email": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Handles for user to modify its emails",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protected",
                    "profile"
                ],
                "summary": "Replaces user's emails",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer + your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "email changes request body",
                        "name": "modifyEmailRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ModifyEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of user's emails with its changes",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Data"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ModifyEmailResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/ptd/profiles/me": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Handles the retrieval of user's full profile data.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protected",
                    "profile"
                ],
                "summary": "Get's the user's profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer + your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "A full profile response of the user",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Data"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.FullProfileResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Error"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/usecase.AppError"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.EmailDetail": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@email.com"
                },
                "isPrimary": {
                    "type": "boolean"
                }
            }
        },
        "dto.FullProfileResponse": {
            "type": "object",
            "properties": {
                "emails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.EmailDetail"
                    }
                },
                "firstName": {
                    "type": "string",
                    "example": "firstName"
                },
                "lastName": {
                    "type": "string",
                    "example": "lastName"
                },
                "phoneNumber": {
                    "type": "string",
                    "example": "0823456786543"
                },
                "username": {
                    "type": "string",
                    "example": "username"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "verystrongpassword"
                },
                "username": {
                    "type": "string",
                    "example": "Username"
                }
            }
        },
        "dto.ModifyEmailRequest": {
            "type": "object",
            "properties": {
                "emails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.EmailDetail"
                    }
                }
            }
        },
        "dto.ModifyEmailResponse": {
            "type": "object",
            "properties": {
                "emails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.EmailDetail"
                    }
                },
                "userId": {
                    "type": "string",
                    "example": "userID"
                }
            }
        },
        "dto.RefreshRequest": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string",
                    "example": "refreshTokenHere"
                }
            }
        },
        "dto.SignUpRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@email.com"
                },
                "firstName": {
                    "type": "string",
                    "example": "FirstName"
                },
                "lastName": {
                    "type": "string",
                    "example": "LastName"
                },
                "password": {
                    "type": "string",
                    "example": "verystrongpassword"
                },
                "phoneNumber": {
                    "type": "string",
                    "example": "628123456789"
                },
                "username": {
                    "type": "string",
                    "example": "Username"
                }
            }
        },
        "dto.TokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.Data": {
            "type": "object",
            "properties": {
                "apiVersion": {
                    "type": "string",
                    "example": "1.0"
                },
                "data": {},
                "paging": {
                    "x-nullable": true,
                    "x-omitempty": true
                },
                "status": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "model.Error": {
            "type": "object",
            "properties": {
                "apiVersion": {
                    "type": "string",
                    "example": "1.0"
                },
                "error": {}
            }
        },
        "usecase.AppError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.AppErrorDetail"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "this is error message"
                }
            }
        },
        "usecase.AppErrorDetail": {
            "type": "object",
            "properties": {
                "domain": {
                    "type": "string",
                    "example": "domain error"
                },
                "message": {
                    "type": "string",
                    "example": "this is descriptive error message"
                },
                "reason": {
                    "type": "string",
                    "example": "this is descriptive error reason"
                },
                "report": {
                    "type": "string",
                    "example": "Please report incident to https://your-report.com"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
