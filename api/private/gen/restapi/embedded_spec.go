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
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userIndexListDefaultBody"
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
              "$ref": "#/definitions/UserStockIndex"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/UserStockIndex"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userIndexSaveDefaultBody"
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
              "$ref": "#/definitions/UserStockIndex"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userIndexRenameDefaultBody"
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
              "$ref": "#/definitions/UserStockIndex"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userIndexGetDefaultBody"
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
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userIndexDeleteDefaultBody"
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
        "operationId": "UserSettingList",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/userSettingListOKBody"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userSettingListDefaultBody"
            }
          }
        }
      },
      "post": {
        "summary": "save",
        "operationId": "UserSettingSave",
        "parameters": [
          {
            "description": "setting",
            "name": "Setting",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserSetting"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/UserSetting"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userSettingSaveDefaultBody"
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
        "operationId": "UserSettingGet",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/UserSetting"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userSettingGetDefaultBody"
            }
          }
        }
      },
      "delete": {
        "operationId": "UserSettingDelete",
        "responses": {
          "200": {
            "description": "ok"
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userSettingDeleteDefaultBody"
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
        "parameters": [
          {
            "type": "string",
            "description": "page token",
            "name": "pageToken",
            "in": "query"
          },
          {
            "type": "string",
            "format": "int32",
            "description": "page size",
            "name": "pageSize",
            "in": "query"
          },
          {
            "type": "string",
            "description": "sort",
            "name": "sort",
            "in": "query"
          },
          {
            "type": "string",
            "description": "not evaluated",
            "name": "notEvaluated",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/userStockEvaluateListOKBody"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userStockEvaluateListDefaultBody"
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
              "$ref": "#/definitions/UserStockEvaluate"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userStockEvaluateGetDefaultBody"
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
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userIndexEvaluateListDefaultBody"
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
              "$ref": "#/definitions/UserIndexEvaluate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/UserIndexEvaluate"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userIndexEvaluateSaveDefaultBody"
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
              "$ref": "#/definitions/UserIndexEvaluate"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/userIndexEvaluateGetDefaultBody"
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
    "UserIndexEvaluate": {
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
    "UserSetting": {
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
    "UserStockEvaluate": {
      "id": "StockEvaluate",
      "description": "stock evaluate",
      "type": "object",
      "properties": {
        "cityNameCN": {
          "description": "City name cn",
          "type": "string"
        },
        "evalRemark": {
          "description": "remark",
          "type": "string"
        },
        "exchangeId": {
          "description": "Exchange id",
          "type": "string"
        },
        "exchangeName": {
          "description": "Exchange name",
          "type": "string"
        },
        "industryName": {
          "description": "Industry name",
          "type": "string"
        },
        "launchDate": {
          "description": "Launch date",
          "type": "string",
          "format": "date-time"
        },
        "provinceNameCN": {
          "description": "Province name cn",
          "type": "string"
        },
        "stockCode": {
          "description": "Stock code",
          "type": "string"
        },
        "stockId": {
          "description": "stock id",
          "type": "string"
        },
        "stockNameCN": {
          "description": "Stock name cn",
          "type": "string"
        },
        "totalScore": {
          "description": "score",
          "type": "string",
          "format": "double"
        },
        "websiteUrl": {
          "description": "Website url",
          "type": "string"
        }
      }
    },
    "UserStockIndex": {
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
    "userIndexDeleteDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexDeleteDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userIndexDeleteDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexEvaluateGetDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexEvaluateGetDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userIndexEvaluateGetDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexEvaluateListDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexEvaluateListDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userIndexEvaluateListDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexEvaluateListOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/UserIndexEvaluate"
      },
      "x-go-gen-location": "operations"
    },
    "userIndexEvaluateSaveDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexEvaluateSaveDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userIndexEvaluateSaveDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexGetDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexGetDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userIndexGetDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexListDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexListDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userIndexListDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexListOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/UserStockIndex"
      },
      "x-go-gen-location": "operations"
    },
    "userIndexRenameDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexRenameDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userIndexRenameDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexSaveDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userIndexSaveDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userIndexSaveDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userSettingDeleteDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userSettingDeleteDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userSettingDeleteDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userSettingGetDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userSettingGetDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userSettingGetDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userSettingListDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userSettingListDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userSettingListDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userSettingListOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/UserSetting"
      },
      "x-go-gen-location": "operations"
    },
    "userSettingSaveDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userSettingSaveDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userSettingSaveDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userStockEvaluateGetDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userStockEvaluateGetDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userStockEvaluateGetDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userStockEvaluateListDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userStockEvaluateListDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "userStockEvaluateListDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "userStockEvaluateListOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/UserStockEvaluate"
      },
      "x-go-gen-location": "operations"
    }
  },
  "responses": {
    "ErrorResponse": {
      "description": "Error response",
      "schema": {
        "type": "object",
        "properties": {
          "code": {
            "description": "Error code",
            "type": "string"
          },
          "errors": {
            "$ref": "#/definitions/userIndexEvaluateGetDefaultBodyErrors"
          },
          "message": {
            "description": "Error message",
            "type": "string"
          },
          "status": {
            "type": "string",
            "format": "int32",
            "default": "Http status"
          }
        }
      }
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
