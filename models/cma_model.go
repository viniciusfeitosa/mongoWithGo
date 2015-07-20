package models

import "gopkg.in/mgo.v2/bson"

// CmaJSON is the struct to persistence of mongo
type CmaJSON struct {
	ID         bson.ObjectId `bson:"_id"`
	Identifier string        `json:"identifier" bson:"identifier"`
	Service    []ServiceJSON `json:"service" bson:"service"`
}

// ServiceJSON is the struct to persistence of mongo
type ServiceJSON struct {
	Active     bool           `json:"active" bson:"active"`
	Categories []CategoryJSON `json:"categories" bson:"categories"`
	Name       string         `json:"name" bson:"name"`
	ServiceID  int            `json:"service_id" bson:"service_id"`
	Slug       string         `json:"slug" bson:"slug"`
}

// CategoryJSON is the struct to persistence of mongo
type CategoryJSON struct {
	Active          bool      `json:"active" bson:"active"`
	Apps            []AppJSON `json:"apps" bson:"apps"`
	Description     string    `json:"description" bson:"description"`
	MessageMultiple string    `json:"message_multiple" bson:"message_multiple"`
	MessageSingle   string    `json:"message_single" bson:"message_single"`
	Name            string    `json:"name" bson:"name"`
	ServiceID       int       `json:"service_id" bson:"service_id"`
	Slug            string    `json:"slug" bson:"slug"`
}

// AppJSON is the struct to persistence of mongo
type AppJSON struct {
	AppID      int    `json:"app_id" bson:"app_id"`
	CategoryID int    `json:"category_id" bson:"category_id"`
	Slug       string `json:"slug" bson:"slug"`
}
