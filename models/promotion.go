package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Promotion -
type Promotion struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	StartDate   string        `json:"eventsdate"`
	EndDate     string        `json:"eventedate"`
	Time        string        `json:"eventtime"`
	Place       string        `json:"place"`
	PostTime    time.Time     `json:"posttime"`
}
