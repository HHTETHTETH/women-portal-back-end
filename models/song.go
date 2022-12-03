package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Singer -
type Singer struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string        `json:"name"`
	Photo string        `json:"photo"`
}

// Song -
type Song struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Singer   Singer        `json:"singer"`
	Album    string        `json:"album"`
	Image    string        `json:"image"`
	MP3      string        `json:"mp3"`
	PostTime time.Time     `json:"posttime"`
}
