// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

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
        "/add-reservation": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add reservation",
                "parameters": [
                    {
                        "description": "Reservation parametres",
                        "name": "Reservation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateReservationDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.NewReservationDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-by-id/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get reservation by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReservationDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-room-reservations/{room_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Search reservation by room id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Room id",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReservationsArrayDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/search-by-phone/{phone}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Search reservation by phone",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Client phone",
                        "name": "phone",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReservationsArrayDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.CreateReservationDto": {
            "type": "object",
            "properties": {
                "clientEmail": {
                    "type": "string"
                },
                "clientFirstName": {
                    "type": "string"
                },
                "clientLastName": {
                    "type": "string"
                },
                "clientPhone": {
                    "type": "string"
                },
                "inTime": {
                    "type": "string"
                },
                "outTime": {
                    "type": "string"
                },
                "roomId": {
                    "type": "string"
                }
            }
        },
        "handlers.NewReservationDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "paymentUrl": {
                    "type": "string"
                }
            }
        },
        "handlers.ReservationDto": {
            "type": "object",
            "properties": {
                "clientEmail": {
                    "type": "string"
                },
                "clientFirstName": {
                    "type": "string"
                },
                "clientLastName": {
                    "type": "string"
                },
                "clientPhone": {
                    "type": "string"
                },
                "cost": {
                    "type": "integer"
                },
                "inTime": {
                    "type": "string"
                },
                "outTime": {
                    "type": "string"
                },
                "roomId": {
                    "type": "string"
                }
            }
        },
        "handlers.ReservationsArrayDto": {
            "type": "object",
            "properties": {
                "reservations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.ReservationDto"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.2.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Booking Service",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
