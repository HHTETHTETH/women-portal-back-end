package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Inspire -
type Inspire struct {
	ID       bson.ObjectId    `bson:"_id,omitempty"`
	Name     string           `json:"name"`
	Skills   string           `json:"skills"`
	Profile  string           `json:"profile"`
	Area     string           `json:"area"`
	Date     string           `json:"date"`
	Images   []InspireImage   `json:"images"`
	History  []InspireHistory `json:"history"`
	PostTime time.Time        `json:"posttime"`
}

// InspireImage -
type InspireImage struct {
	Src string `json:"src"`
}

// InspireHistory -
type InspireHistory struct {
	Text     string `json:"text"`
	Title    string `json:"title"`
	ImageURL string `json:"imageUrl"`
}
