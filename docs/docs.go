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
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Courses"
                ],
                "summary": "Get all courses",
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
                "description": "Get all golfs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Golfs"
                ],
                "summary": "Get all golfs",
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
                "golf_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "num_holes": {
                    "type": "integer"
                },
                "pitch_and_putt": {
                    "type": "boolean"
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
                    "example": "Monterey"
                },
                "googleMapLinks": {
                    "description": "@Description Google Maps link for the golf course",
                    "type": "string",
                    "example": "https://maps.google.com/?q=Pebble+Beach"
                },
                "id": {
                    "type": "string"
                },
                "latitude": {
                    "description": "@Description Latitude of the golf course",
                    "type": "number",
                    "example": 36.567
                },
                "longitude": {
                    "description": "@Description Longitude of the golf course",
                    "type": "number",
                    "example": -121.95
                },
                "name": {
                    "description": "@Description Name of the golf course",
                    "type": "string",
                    "example": "Pebble Beach"
                },
                "postalCode": {
                    "description": "@Description Postal code of the golf course location",
                    "type": "string",
                    "example": "93953"
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
                "holeNumber": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "par": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
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
