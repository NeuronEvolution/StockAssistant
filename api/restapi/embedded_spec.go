// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json;charset=utf-8"
  ],
  "produces": [
    "application/json;charset=utf-8"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Stock Assistant",
    "contact": {
      "name": "mars"
    },
    "version": "v1"
  },
  "basePath": "/api/stock-assistant/v1",
  "paths": {
    "/{userId}/indices": {
      "get": {
        "summary": "Get user indices",
        "operationId": "UserIndexList",
        "security": [
          {}
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/userIndexListOKBody"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "summary": "save",
        "operationId": "UserIndexSave",
        "parameters": [
          {
            "description": "Index",
            "name": "Index",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/StockIndex"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/StockIndex"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "User id",
          "name": "userId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{userId}/indices/rename": {
      "post": {
        "operationId": "UserIndexRename",
        "parameters": [
          {
            "type": "string",
            "description": "old name",
            "name": "oldName",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "new name",
            "name": "newName",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/StockIndex"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "User id",
          "name": "userId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{userId}/indices/{indexId}": {
      "get": {
        "summary": "Get user index",
        "operationId": "UserIndexGet",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/StockIndex"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "operationId": "UserIndexDelete",
        "responses": {
          "200": {
            "description": "ok"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "User id",
          "name": "userId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "index id",
          "name": "indexId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{userId}/settings": {
      "get": {
        "summary": "list",
        "operationId": "UserSettingsList",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/userSettingsListOKBody"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "summary": "save",
        "operationId": "UserSettingsSave",
        "parameters": [
          {
            "description": "setting",
            "name": "Setting",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Setting"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/Setting"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "User id",
          "name": "userId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{userId}/settings/{configKey}": {
      "get": {
        "operationId": "UserSettingsGet",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/Setting"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "operationId": "UserSettingsDelete",
        "responses": {
          "200": {
            "description": "ok"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "User id",
          "name": "userId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "config key",
          "name": "configKey",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{userId}/stockEvaluates": {
      "get": {
        "operationId": "UserStockEvaluateList",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/userStockEvaluateListOKBody"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "summary": "save",
        "operationId": "UserStockEvaluateSave",
        "parameters": [
          {
            "name": "stockEvaluate",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/StockEvaluate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/StockEvaluate"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "User id",
          "name": "userId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{userId}/stockEvaluates/{stockId}": {
      "get": {
        "operationId": "UserStockEvaluateGet",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/StockEvaluate"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "User id",
          "name": "userId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "stock id",
          "name": "stockId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{userId}/stockEvaluates/{stockId}/indexEvaluates": {
      "get": {
        "operationId": "UserIndexEvaluateList",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/userIndexEvaluateListOKBody"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "operationId": "UserIndexEvaluateSave",
        "parameters": [
          {
            "name": "indexEvaluate",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/IndexEvaluate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/IndexEvaluate"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "User id",
          "name": "userId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "stock id",
          "name": "stockId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{userId}/stockEvaluates/{stockId}/indexEvaluates/{indexName}": {
      "get": {
        "operationId": "UserIndexEvaluateGet",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/IndexEvaluate"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "User id",
          "name": "userId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "stock id",
          "name": "stockId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "description": "index name",
          "name": "indexName",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "IndexEvaluate": {
      "id": "IndexEvaluate",
      "description": "index evaluate",
      "type": "object",
      "properties": {
        "evalRemark": {
          "description": "eval remark",
          "type": "string"
        },
        "evalStars": {
          "description": "eval stars",
          "type": "string",
          "format": "int32"
        },
        "indexName": {
          "description": "index name",
          "type": "string"
        },
        "updateTime": {
          "description": "update time",
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "Setting": {
      "id": "Setting",
      "description": "User stock setting",
      "type": "object",
      "properties": {
        "key": {
          "description": "Key",
          "type": "string"
        },
        "value": {
          "description": "Value",
          "type": "string"
        }
      }
    },
    "StockEvaluate": {
      "id": "StockEvaluate",
      "description": "stock evaluate",
      "type": "object",
      "properties": {
        "evalRemark": {
          "description": "remark",
          "type": "string"
        },
        "stockId": {
          "description": "stock id",
          "type": "string"
        },
        "totalScore": {
          "description": "score",
          "type": "string",
          "format": "double"
        }
      }
    },
    "StockIndex": {
      "id": "Index",
      "description": "User stock index",
      "type": "object",
      "properties": {
        "aiWeight": {
          "description": "ai weight",
          "type": "string",
          "format": "int32"
        },
        "desc": {
          "description": "desc",
          "type": "string"
        },
        "evalWeight": {
          "description": "Eval weight",
          "type": "string",
          "format": "int32"
        },
        "name": {
          "description": "name",
          "type": "string"
        },
        "niWeight": {
          "description": "ni weight",
          "type": "string",
          "format": "int32"
        }
      }
    },
    "userIndexEvaluateListOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/IndexEvaluate"
      },
      "x-go-gen-location": "operations"
    },
    "userIndexListOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/StockIndex"
      },
      "x-go-gen-location": "operations"
    },
    "userSettingsListOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Setting"
      },
      "x-go-gen-location": "operations"
    },
    "userStockEvaluateListOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/StockEvaluate"
      },
      "x-go-gen-location": "operations"
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
