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
        "/v1/health/": {
            "get": {
                "description": "Health Check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/validate": {
            "post": {
                "description": "Validate User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Validate User",
                "parameters": [
                    {
                        "description": "ValidateUserRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.ValidateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_itera-io_openstack-web-client_api_dto.ValidateUserRequest": {
            "type": "object",
            "required": [
                "password",
                "url",
                "username"
            ],
            "properties": {
                "applicationCredEnabled": {
                    "type": "boolean"
                },
                "domain": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "url": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "result": {},
                "resultCode": {
                    "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.ResultCode"
                },
                "success": {
                    "type": "boolean"
                },
                "validationErrors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_validations.ValidationError"
                    }
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_helper.ResultCode": {
            "type": "integer",
            "enum": [
                0,
                40001,
                40101,
                40301,
                40401,
                42901,
                42902,
                50001,
                50002
            ],
            "x-enum-varnames": [
                "Success",
                "ValidationError",
                "AuthError",
                "ForbiddenError",
                "NotFoundError",
                "LimiterError",
                "OtpLimiterError",
                "CustomRecovery",
                "InternalError"
            ]
        },
        "github_com_itera-io_openstack-web-client_api_validations.ValidationError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "property": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
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
