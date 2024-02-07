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
        "/opengate/config/": {
            "get": {
                "description": "Retrieve a list of all service configurations.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config"
                ],
                "summary": "Get all service configurations",
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.ListConfigResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            },
            "put": {
                "description": "Create or update service configuration based on the provided request.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config"
                ],
                "summary": "Create or update service configuration.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ping ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body containing configuration details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateServiceConfigRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateServiceConfigResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            }
        },
        "/opengate/config/{id}": {
            "get": {
                "description": "Retrieve a service configuration based on the provided request containing the ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config"
                ],
                "summary": "Get a service configuration by ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the service configuration",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/dto.ConfigByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "404": {
                        "description": "Config not found",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a service configuration based on the provided ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config"
                ],
                "summary": "Delete a service configuration by ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the service configuration to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteConfigResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid ID format",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "404": {
                        "description": "Config not found",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.CustomError"
                        }
                    }
                }
            }
        },
        "/opengate/ping/": {
            "get": {
                "description": "Pings the server and returns \"Okay\" if successful.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ping"
                ],
                "summary": "Pings the server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.PingResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ConfigByIdResponse": {
            "type": "object",
            "properties": {
                "config": {
                    "$ref": "#/definitions/dto.ServiceConfig"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateServiceConfigRequest": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "endpoint": {
                    "type": "string"
                },
                "regex": {
                    "type": "string"
                }
            }
        },
        "dto.CreateServiceConfigResponse": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "dto.DeleteConfigResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "dto.ListConfigResponse": {
            "type": "object",
            "properties": {
                "configs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ServiceConfig"
                    }
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "dto.PingResponse": {
            "type": "object",
            "properties": {
                "message": {},
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "dto.ServiceConfig": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "createdOn": {
                    "type": "string"
                },
                "endpoint": {
                    "type": "string"
                },
                "regex": {
                    "type": "string"
                },
                "updatedOn": {
                    "type": "string"
                }
            }
        },
        "utils.CustomError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
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
