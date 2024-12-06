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
        "/api/health": {
            "get": {
                "description": "Get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System Management"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/images": {
            "get": {
                "description": "Get a list of docker image",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Docker Image Management"
                ],
                "summary": "List of Docker Image",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/images.Image"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Pulls a docker image and register it into the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Docker Image Management"
                ],
                "summary": "Pull Docker Image",
                "parameters": [
                    {
                        "description": "Image Input",
                        "name": "image",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/images.PullImageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/images.PullImageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/images/{id}": {
            "get": {
                "description": "Get details about a docker image",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Docker Image Management"
                ],
                "summary": "Details about a Docker Image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Image ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/images.Image"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Removes a docker image form the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Docker Image Management"
                ],
                "summary": "Remove Docker Image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Image ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/images.ImageDeleteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/sandboxes": {
            "get": {
                "description": "List sandbox environments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sandbox Management"
                ],
                "summary": "List sandbox environments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ContainerInfo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.ContainerInfo": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "images.Image": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2013-08-20T18:52:09.000Z"
                },
                "id": {
                    "type": "string",
                    "example": "a407dee395ed"
                },
                "image_name": {
                    "type": "string",
                    "example": "dockware/dev"
                },
                "image_tag": {
                    "type": "string",
                    "example": "6.6.8.2"
                },
                "size": {
                    "type": "integer",
                    "example": 1048576
                }
            }
        },
        "images.ImageDeleteResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Docker Image removed successfully"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "images.PullImageRequest": {
            "type": "object",
            "properties": {
                "image_name": {
                    "type": "string",
                    "example": "dockware/dev"
                },
                "image_tag": {
                    "type": "string",
                    "example": "6.6.8.2"
                }
            }
        },
        "images.PullImageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Docker Image created successfully"
                },
                "status": {
                    "type": "string",
                    "example": "success"
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
