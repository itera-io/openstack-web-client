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
        "/v1/users/auth": {
            "post": {
                "description": "Authenticate User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Authenticate User",
                "parameters": [
                    {
                        "description": "AuthenticateUserRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.AuthenticateUserRequest"
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
        },
        "/v2/flavors": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "List Flavor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Flavors"
                ],
                "summary": "List Flavor",
                "parameters": [
                    {
                        "description": "ListFlavor Request",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.ListFlavorRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ListFlavor response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.ListFlavorResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v2/flavors/{id}": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Get Flavor by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Flavors"
                ],
                "summary": "Get Flavor",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetFlavor response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.GetFlavorResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v2/images": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "List Image",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "List Image",
                "parameters": [
                    {
                        "description": "ListImage Request",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.ListImageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ListImage response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.ListImageResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v2/images/{id}": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Get Image by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "Get Image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetImage response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.GetImageResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v3/Projects": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "List Project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "List Project",
                "parameters": [
                    {
                        "description": "ListProjectRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.ListProjectRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ListProject response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.ListProjectResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v3/regions": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "List Region",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Regions"
                ],
                "summary": "List Region",
                "parameters": [
                    {
                        "description": "ListRegionRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.ListRegionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ListRegion response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.ListRegionResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v3/users/{id}/projects": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "List UserProject",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserProjects"
                ],
                "summary": "List UserProject",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ListUserProject response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.ListUserProjectResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized request",
                        "schema": {
                            "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_itera-io_openstack-web-client_api_dto.AuthenticateUserRequest": {
            "type": "object",
            "required": [
                "password",
                "url",
                "username"
            ],
            "properties": {
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
        "github_com_itera-io_openstack-web-client_api_dto.CommonDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_dto.FlavorDto": {
            "type": "object",
            "properties": {
                "disk": {
                    "description": "Disk is the amount of root disk, measured in GB.",
                    "type": "integer"
                },
                "id": {
                    "description": "ID is the flavor's unique ID.",
                    "type": "string"
                },
                "name": {
                    "description": "Name is the name of the flavor.",
                    "type": "string"
                },
                "os-flavor-access:is_public": {
                    "description": "IsPublic indicates whether the flavor is public.",
                    "type": "boolean"
                },
                "ram": {
                    "description": "RAM is the amount of memory, measured in MB.",
                    "type": "integer"
                },
                "vcpus": {
                    "description": "VCPUs indicates how many (virtual) CPUs are available for this flavor.",
                    "type": "integer"
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_dto.GetFlavorResponse": {
            "type": "object",
            "properties": {
                "flavor": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.FlavorDto"
                    }
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_dto.GetImageResponse": {
            "type": "object",
            "properties": {
                "image": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.ImageDto"
                    }
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_dto.ImageDto": {
            "type": "object",
            "properties": {
                "disk": {
                    "description": "Disk is the amount of root disk, measured in GB.",
                    "type": "integer"
                },
                "id": {
                    "description": "ID is the flavor's unique ID.",
                    "type": "string"
                },
                "name": {
                    "description": "Name is the name of the flavor.",
                    "type": "string"
                },
                "ram": {
                    "description": "RAM is the amount of memory, measured in MB.",
                    "type": "integer"
                },
                "status": {
                    "description": "VCPUs indicates how many (virtual) CPUs are available for this flavor.",
                    "type": "string"
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_dto.ListFlavorRequest": {
            "type": "object",
            "properties": {
                "minDisk": {
                    "type": "integer"
                },
                "minRam": {
                    "type": "integer"
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_dto.ListFlavorResponse": {
            "type": "object",
            "properties": {
                "flavors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.FlavorDto"
                    }
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_dto.ListImageRequest": {
            "type": "object"
        },
        "github_com_itera-io_openstack-web-client_api_dto.ListImageResponse": {
            "type": "object",
            "properties": {
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.CommonDto"
                    }
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_dto.ListProjectRequest": {
            "type": "object"
        },
        "github_com_itera-io_openstack-web-client_api_dto.ListProjectResponse": {
            "type": "object",
            "properties": {
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.CommonDto"
                    }
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_dto.ListRegionRequest": {
            "type": "object"
        },
        "github_com_itera-io_openstack-web-client_api_dto.ListRegionResponse": {
            "type": "object",
            "properties": {
                "regions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.CommonDto"
                    }
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_dto.ListUserProjectResponse": {
            "type": "object",
            "properties": {
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_itera-io_openstack-web-client_api_dto.CommonDto"
                    }
                }
            }
        },
        "github_com_itera-io_openstack-web-client_api_dto.ValidateUserRequest": {
            "type": "object",
            "required": [
                "password",
                "url",
                "username"
            ],
            "properties": {
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
