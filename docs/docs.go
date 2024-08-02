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
        "/v1/courses": {
            "get": {
                "description": "Get all courses with pagination",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Courses"
                ],
                "summary": "Get all courses",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number (default is 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size (default is 10)",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Course"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Courses"
                ],
                "summary": "Create a new course",
                "parameters": [
                    {
                        "description": "Course",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Course"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/Course"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            }
        },
        "/v1/courses/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Courses"
                ],
                "summary": "Get a course by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Course"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Courses"
                ],
                "summary": "Delete a course by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Courses"
                ],
                "summary": "Update a course by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Course",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Course"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Course"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            }
        },
        "/v1/golfs": {
            "get": {
                "description": "Get all golfs with pagination",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Golfs"
                ],
                "summary": "Get all golfs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number (default is 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size (default is 10)",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Golf"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new golf",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Golfs"
                ],
                "summary": "Create a new golf",
                "parameters": [
                    {
                        "description": "Golf",
                        "name": "golf",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Golf"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/Golf"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            }
        },
        "/v1/golfs/{golf_id}/courses": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Courses"
                ],
                "summary": "Get all courses for a specific golf",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Golf ID",
                        "name": "golf_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Course"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            }
        },
        "/v1/golfs/{id}": {
            "get": {
                "description": "Get a golf by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Golfs"
                ],
                "summary": "Get a golf by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Golf ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Golf"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an existing golf by ID",
                "tags": [
                    "Golfs"
                ],
                "summary": "Delete an existing golf",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Golf ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update an existing golf by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Golfs"
                ],
                "summary": "Update an existing golf",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Golf ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Golf",
                        "name": "golf",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Golf"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Golf"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            }
        },
        "/v1/holes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Holes"
                ],
                "summary": "Get all holes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Hole"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Holes"
                ],
                "summary": "Create a new hole",
                "parameters": [
                    {
                        "description": "Hole",
                        "name": "hole",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Hole"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/Hole"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            }
        },
        "/v1/holes/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Holes"
                ],
                "summary": "Get a hole by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hole ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Hole"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Holes"
                ],
                "summary": "Delete an existing hole",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hole ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Holes"
                ],
                "summary": "Update an existing hole",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hole ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Hole",
                        "name": "hole",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Hole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Hole"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            }
        },
        "/v1/tees": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tees"
                ],
                "summary": "Get all tees",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Tee"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tees"
                ],
                "summary": "Create a new tee",
                "parameters": [
                    {
                        "description": "Tee",
                        "name": "tee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Tee"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/Tee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            }
        },
        "/v1/tees/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tees"
                ],
                "summary": "Get a tee by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Tee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Tees"
                ],
                "summary": "Delete an existing tee",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tees"
                ],
                "summary": "Update an existing tee",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Tee",
                        "name": "tee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Tee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Tee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "AppError": {
            "description": "Error that can be returned by the application",
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "Course": {
            "description": "Model for a golf course",
            "type": "object",
            "properties": {
                "compact": {
                    "type": "boolean"
                },
                "createdAt": {
                    "description": "@Description Creation date of the model, automatically generated",
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z (auto-generated)"
                },
                "golf_id": {
                    "type": "string"
                },
                "id": {
                    "description": "@Description ID of the model, automatically generated",
                    "type": "string",
                    "example": "123e4567-e89b-12d3-a456-426614174000 (auto-generated)"
                },
                "name": {
                    "type": "string"
                },
                "num_holes": {
                    "type": "integer"
                },
                "pitch_and_putt": {
                    "type": "boolean"
                },
                "updatedAt": {
                    "description": "@Description Last update date of the model, automatically generated",
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z (auto-generated)"
                }
            }
        },
        "Golf": {
            "description": "Model for a golf",
            "type": "object",
            "properties": {
                "city": {
                    "description": "@Description City where the golf course is located",
                    "type": "string",
                    "example": "Guyancourt"
                },
                "createdAt": {
                    "description": "@Description Creation date of the model, automatically generated",
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z (auto-generated)"
                },
                "googleMapLinks": {
                    "description": "@Description Google Maps link for the golf course",
                    "type": "string",
                    "example": "https://maps.google.com/?q=Golf+National"
                },
                "id": {
                    "description": "@Description ID of the model, automatically generated",
                    "type": "string",
                    "example": "123e4567-e89b-12d3-a456-426614174000 (auto-generated)"
                },
                "latitude": {
                    "description": "@Description Latitude of the golf course",
                    "type": "number",
                    "example": 48.754
                },
                "longitude": {
                    "description": "@Description Longitude of the golf course",
                    "type": "number",
                    "example": 2.074
                },
                "name": {
                    "description": "@Description Name of the golf course",
                    "type": "string",
                    "example": "Golf National"
                },
                "postalCode": {
                    "description": "@Description Postal code of the golf course location",
                    "type": "string",
                    "example": "78280"
                },
                "updatedAt": {
                    "description": "@Description Last update date of the model, automatically generated",
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z (auto-generated)"
                }
            }
        },
        "Hole": {
            "description": "Model for a hole",
            "type": "object",
            "properties": {
                "courseID": {
                    "type": "string"
                },
                "createdAt": {
                    "description": "@Description Creation date of the model, automatically generated",
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z (auto-generated)"
                },
                "holeNumber": {
                    "type": "integer"
                },
                "id": {
                    "description": "@Description ID of the model, automatically generated",
                    "type": "string",
                    "example": "123e4567-e89b-12d3-a456-426614174000 (auto-generated)"
                },
                "par": {
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "@Description Last update date of the model, automatically generated",
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z (auto-generated)"
                }
            }
        },
        "Tee": {
            "description": "Model for a tee",
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "createdAt": {
                    "description": "@Description Creation date of the model, automatically generated",
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z (auto-generated)"
                },
                "distance": {
                    "type": "integer"
                },
                "holeID": {
                    "type": "string"
                },
                "id": {
                    "description": "@Description ID of the model, automatically generated",
                    "type": "string",
                    "example": "123e4567-e89b-12d3-a456-426614174000 (auto-generated)"
                },
                "updatedAt": {
                    "description": "@Description Last update date of the model, automatically generated",
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z (auto-generated)"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0-Beta",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "ProjectPAR API",
	Description:      "This project is a comprehensive golf management system designed to facilitate the management of golf courses, including tracking of golf courses, holes, and tees.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
