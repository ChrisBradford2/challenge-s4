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
        "/hackathons": {
            "get": {
                "description": "Récupère une liste de tous les hackathons",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hackathons"
                ],
                "summary": "Lire tous les hackathons",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Hackathon"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Ajoute un nouveau hackathon à la base de données",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hackathons"
                ],
                "summary": "Créer un hackathon",
                "parameters": [
                    {
                        "description": "Hackathon à créer",
                        "name": "hackathon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.HackathonCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Hackathon"
                        }
                    }
                }
            }
        },
        "/hackathons/{id}": {
            "get": {
                "description": "Récupère un hackathon par son ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hackathons"
                ],
                "summary": "Lire un hackathon spécifique",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID du Hackathon",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Hackathon"
                        }
                    }
                }
            },
            "put": {
                "description": "Met à jour les informations d'un hackathon par son ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hackathons"
                ],
                "summary": "Mettre à jour un hackathon",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID du Hackathon",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Informations du Hackathon à mettre à jour",
                        "name": "hackathon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Hackathon"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Hackathon"
                        }
                    }
                }
            },
            "delete": {
                "description": "Supprime un hackathon par son ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hackathons"
                ],
                "summary": "Supprimer un hackathon",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID du Hackathon",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "true si la suppression est réussie",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Hackathon": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "maxParticipants": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.HackathonCreate": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Un événement pour les développeurs"
                },
                "location": {
                    "type": "string",
                    "example": "Paris"
                },
                "maxParticipants": {
                    "type": "integer",
                    "example": 100
                },
                "name": {
                    "type": "string",
                    "example": "Hackathon de Paris"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Kiwi Collective API",
	Description:      "Swagger API for the Kiwi Collective project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
