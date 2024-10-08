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
        "/task/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "CreateTask  users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "CreateTask",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/tasksentity.CreateTaskReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/tasksentity.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/delete/{task_id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "DeleteTask tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "DeleteTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task_id",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/tasks": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetTasks tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "GetTasks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Field",
                        "name": "field",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Value",
                        "name": "value",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/tasksentity.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/update/{task_id}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "UpdateTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task_id",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/tasksentity.UpdateTaskReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/tasksentity.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/{field}/{value}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetTask task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "GetTask",
                "parameters": [
                    {
                        "type": "string",
                        "description": "field",
                        "name": "field",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "value",
                        "name": "value",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/tasksentity.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Authenticate a user and return a token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User login information",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entityuser.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entityuser.LoginReq"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "User registration information",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entityuser.CreateUserReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entityuser.StatusMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/user/verify": {
            "post": {
                "description": "Confirm the code sent to the email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "VerifyUser",
                "parameters": [
                    {
                        "description": "User verification information",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entityuser.VerifyUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entityuser.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entityuser.CreateUserReq": {
            "type": "object",
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "entityuser.LoginReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "entityuser.StatusMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "entityuser.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "entityuser.VerifyUserReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "secret_code": {
                    "type": "string"
                }
            }
        },
        "tasksentity.CreateTaskReq": {
            "type": "object",
            "properties": {
                "status": {
                    "$ref": "#/definitions/tasksentity.Status"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "tasksentity.Status": {
            "type": "object",
            "properties": {
                "condition": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "important": {
                    "type": "boolean"
                }
            }
        },
        "tasksentity.Task": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/tasksentity.Status"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "tasksentity.UpdateTaskReq": {
            "type": "object",
            "properties": {
                "status": {
                    "$ref": "#/definitions/tasksentity.Status"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "diyorbek.touristan-bs.uz:9000",
	BasePath:         "/",
	Schemes:          []string{"https"},
	Title:            "Artisan Connect",
	Description:      "This is a sample server for a restaurant reservation system.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
