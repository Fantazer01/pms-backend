// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Tokens"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/projects": {
            "get": {
                "description": "Get projects",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Get projects",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Page index of projects",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Page size of projects",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProjectsPaged"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            },
            "post": {
                "description": "Create project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Create project",
                "parameters": [
                    {
                        "description": "Project",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.InsertProject"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Project"
                        }
                    },
                    "422": {
                        "description": "Bind error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/projects/archived": {
            "get": {
                "description": "Get archived projects",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Get archived projects",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Page index of projects",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Page size of projects",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProjectsPaged"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}": {
            "get": {
                "description": "Get project by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Get project by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Project"
                        }
                    },
                    "404": {
                        "description": "Project not found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "422": {
                        "description": "Incorrect id of project",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            },
            "put": {
                "description": "Update project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Update project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Project",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.InsertProject"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Project"
                        }
                    },
                    "404": {
                        "description": "Project not found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "422": {
                        "description": "Incorrect id of project/Bind error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Delete project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Project not found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "422": {
                        "description": "Incorrect id of project",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}/archive": {
            "put": {
                "description": "Archive the project by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Archive the project by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProjectsPaged"
                        }
                    },
                    "404": {
                        "description": "Project not found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "422": {
                        "description": "Incorrect id of project",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}/members": {
            "get": {
                "description": "Get project members",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Get project members",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UsersPaged"
                        }
                    },
                    "404": {
                        "description": "Project not found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "422": {
                        "description": "Incorrect id of project",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}/members/{user_id}": {
            "post": {
                "description": "Add project member",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Add project member",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Project not found/User not found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "422": {
                        "description": "Incorrect id of project/Incorrect id of user",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete project member",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Delete project member",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Project not found/User not found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "422": {
                        "description": "Incorrect id of project/Incorrect id of user",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}/tasks": {
            "get": {
                "description": "Get project tasks",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Get project tasks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Tasks"
                        }
                    },
                    "404": {
                        "description": "Project not found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "422": {
                        "description": "Incorrect id of project",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}/unarchive": {
            "put": {
                "description": "Unarchive the project by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Unarchive the project by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProjectsPaged"
                        }
                    },
                    "404": {
                        "description": "Project not found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "422": {
                        "description": "Incorrect id of project",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "description": "Refresh",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Tokens"
                        }
                    }
                }
            }
        },
        "/task": {
            "post": {
                "description": "Create task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Create task",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                }
            }
        },
        "/task/{id}": {
            "get": {
                "description": "Get task by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Get task by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                }
            },
            "put": {
                "description": "Update task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Update task",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Delete task",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page index",
                        "name": "pageIndex",
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
                            "$ref": "#/definitions/model.UsersPaged"
                        }
                    }
                }
            },
            "post": {
                "description": "Create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create user",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            },
            "put": {
                "description": "Update user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "delete": {
                "description": "Delete User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user",
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.InsertProject": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Project": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.ProjectsPaged": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Project"
                    }
                },
                "page_index": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.Task": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "project_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.Tasks": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Task"
                    }
                }
            }
        },
        "model.Tokens": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "middle_name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.UsersPaged": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "page_index": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total": {
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
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "PMS API",
	Description:      "PMS - project management system. It is course work of student team for 5 course in MEPhI.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
