// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "株式会社フルネス",
            "url": "https://www.fullness.co.jp",
            "email": "furukawa@fullness.co.jp"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/category/add": {
            "post": {
                "description": "商品カテゴリを登録する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品カテゴリを登録する",
                "operationId": "add-category",
                "parameters": [
                    {
                        "description": "name(カテゴリ名)",
                        "name": "newcategory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.CategoryUpResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    }
                }
            }
        },
        "/category/byid/{id}": {
            "get": {
                "description": "指定されたカテゴリ番号の商品カテゴリを取得する",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品カテゴリを取得する",
                "operationId": "get-category-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id(カテゴリ番号)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.Category"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    }
                }
            }
        },
        "/category/delete": {
            "delete": {
                "description": "商品カテゴリを削除する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品カテゴリを削除する",
                "operationId": "delete-category",
                "parameters": [
                    {
                        "description": "カテゴリ番号,カテゴリ名",
                        "name": "newcategory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.CategoryUpResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    }
                }
            }
        },
        "/category/list": {
            "get": {
                "description": "すべての商品カテゴリを取得する",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品カテゴリを取得する",
                "operationId": "get-categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.Category"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    }
                }
            }
        },
        "/category/update": {
            "put": {
                "description": "商品カテゴリを変更する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品カテゴリを変更する",
                "operationId": "update-category",
                "parameters": [
                    {
                        "description": "カテゴリ番号,カテゴリ名",
                        "name": "newcategory",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.CategoryUpResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    }
                }
            }
        },
        "/product/add": {
            "post": {
                "description": "商品を登録する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品を登録する",
                "operationId": "add-product",
                "parameters": [
                    {
                        "description": "name(商品名),price(単価),categoryId(商品カテゴリ番号)",
                        "name": "newproduct",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ProductDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.ProductUpResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    }
                }
            }
        },
        "/product/byid/{id}": {
            "get": {
                "description": "指定された商品番号の商品を取得する",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品を取得する",
                "operationId": "get-product-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id(商品番号)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.Product"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    }
                }
            }
        },
        "/product/bykeyword/{keyword}": {
            "get": {
                "description": "指定されたキーワードが含まれる商品を取得する",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品を取得する",
                "operationId": "get-product-by-keyword",
                "parameters": [
                    {
                        "type": "string",
                        "description": "キーワード",
                        "name": "keyword",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.Product"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    }
                }
            }
        },
        "/product/delete": {
            "delete": {
                "description": "商品を削除する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品を削除する",
                "operationId": "delete-product",
                "parameters": [
                    {
                        "description": "商品番号",
                        "name": "deleteproduct",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ProductDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.ProductUpResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    }
                }
            }
        },
        "/product/list": {
            "get": {
                "description": "すべての商品を取得する",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品を取得する",
                "operationId": "get-products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.Product"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    }
                }
            }
        },
        "/product/update": {
            "put": {
                "description": "商品を変更する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "商品を変更する",
                "operationId": "update-product",
                "parameters": [
                    {
                        "description": "id(商品番号),name(商品名),price(単価),categoryId(商品カテゴリ番号)",
                        "name": "updateproduct",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ProductDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.ProductUpResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CategoryDTO": {
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
        "dto.ProductDTO": {
            "type": "object",
            "properties": {
                "categoryId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "pb.Category": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "商品カテゴリId",
                    "type": "string"
                },
                "name": {
                    "description": "商品カテゴリ名",
                    "type": "string"
                }
            }
        },
        "pb.CategoryUpParam": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "商品カテゴリ番号",
                    "type": "string"
                },
                "name": {
                    "description": "商品カテゴリ名",
                    "type": "string"
                }
            }
        },
        "pb.CategoryUpResult": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "エラー",
                    "allOf": [
                        {
                            "$ref": "#/definitions/pb.Error"
                        }
                    ]
                },
                "param": {
                    "description": "リクエストパラメータ",
                    "allOf": [
                        {
                            "$ref": "#/definitions/pb.CategoryUpParam"
                        }
                    ]
                }
            }
        },
        "pb.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "pb.Product": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "商品Id",
                    "type": "string"
                },
                "name": {
                    "description": "商品名",
                    "type": "string"
                },
                "price": {
                    "description": "単価",
                    "type": "integer"
                }
            }
        },
        "pb.ProductUpParam": {
            "type": "object",
            "properties": {
                "category_id": {
                    "description": "カテゴリ番号",
                    "type": "string"
                },
                "id": {
                    "description": "商品番号",
                    "type": "string"
                },
                "name": {
                    "description": "商品名",
                    "type": "string"
                },
                "price": {
                    "description": "単価",
                    "type": "integer"
                }
            }
        },
        "pb.ProductUpResult": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "エラー",
                    "allOf": [
                        {
                            "$ref": "#/definitions/pb.Error"
                        }
                    ]
                },
                "param": {
                    "description": "リクエストパラメータ",
                    "allOf": [
                        {
                            "$ref": "#/definitions/pb.ProductUpParam"
                        }
                    ]
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "gRPCサンプルプログラム(CQRSサービス)",
	Description:      "商品カテゴリ情報と商品情報を管理するAPIサービス",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
