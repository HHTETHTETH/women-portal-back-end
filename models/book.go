package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Book -
type Book struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Author      string        `json:"author"`
	Cover       string        `json:"cover"`
	Release     string        `json:"release"`
	Description string        `json:"description"`
	PostTime    time.Time     `json:"posttime"`
}
