package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Diy -
type Diy struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Img          string        `json:"img"`
	Video        string        `json:"video"`
	Title        string        `json:"title"`
	Author       string        `json:"author"`
	Date         string        `json:"date"`
	Requirements []string      `json:"requirements"`
	Description  string        `json:"description"`
	PostTime     time.Time     `json:"posttime"`
}
