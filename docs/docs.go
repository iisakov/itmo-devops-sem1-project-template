// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Artisan",
            "url": "http://www.by_artisan.io/support",
            "email": "by@artisan.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v0/price/": {
            "get": {
                "description": "Метод позволяет получить все товары",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Получить все товары",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/model.Item"
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Метод позволяет записать в базу данных список продуктов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Добавить продукт",
                "parameters": [
                    {
                        "description": "Продукты",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Item"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Item": {
            "type": "object",
            "properties": {
                "total_categories": {
                    "type": "integer"
                },
                "total_items": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "price archive [by_Artisan]",
	Description:      "price archive позволяет хранить информацию о ценах на товары.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
