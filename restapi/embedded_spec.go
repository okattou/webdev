// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API DE CREATION",
    "title": "WEB-DEV",
    "version": "1.0.0"
  },
  "host": "api.example.com",
  "basePath": "/v2",
  "paths": {
    "/users": {
      "get": {
        "description": "api wed-dev",
        "produces": [
          "application/yml"
        ],
        "summary": "infos sur les utilisateurs de la base de donnees",
        "responses": {
          "200": {
            "description": "ras"
          }
        }
      },
      "put": {
        "description": "api wed-dev",
        "produces": [
          "application/yml"
        ],
        "summary": "infos sur les utilisateurs de la base de donnees",
        "responses": {
          "200": {
            "description": "ras"
          }
        }
      },
      "post": {
        "description": "api wed-dev",
        "produces": [
          "application/yml"
        ],
        "summary": "infos sur les utilisateurs de la base de donnees",
        "responses": {
          "200": {
            "description": "ras"
          }
        }
      },
      "delete": {
        "description": "api wed-dev",
        "produces": [
          "application/yml"
        ],
        "summary": "infos sur les utilisateurs de la base de donnees",
        "responses": {
          "200": {
            "description": "ras"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API DE CREATION",
    "title": "WEB-DEV",
    "version": "1.0.0"
  },
  "host": "api.example.com",
  "basePath": "/v2",
  "paths": {
    "/users": {
      "get": {
        "description": "api wed-dev",
        "produces": [
          "application/yml"
        ],
        "summary": "infos sur les utilisateurs de la base de donnees",
        "responses": {
          "200": {
            "description": "ras"
          }
        }
      },
      "put": {
        "description": "api wed-dev",
        "produces": [
          "application/yml"
        ],
        "summary": "infos sur les utilisateurs de la base de donnees",
        "responses": {
          "200": {
            "description": "ras"
          }
        }
      },
      "post": {
        "description": "api wed-dev",
        "produces": [
          "application/yml"
        ],
        "summary": "infos sur les utilisateurs de la base de donnees",
        "responses": {
          "200": {
            "description": "ras"
          }
        }
      },
      "delete": {
        "description": "api wed-dev",
        "produces": [
          "application/yml"
        ],
        "summary": "infos sur les utilisateurs de la base de donnees",
        "responses": {
          "200": {
            "description": "ras"
          }
        }
      }
    }
  }
}`))
}
