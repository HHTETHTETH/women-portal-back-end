package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Fashion -
type Fashion struct {
	ID          bson.ObjectId    `bson:"_id,omitempty"`
	Album       string           `json:"album"`
	Date        string           `json:"date"`
	Description string           `json:"description"`
	Gallery     []FashionGallery `json:"gallery"`
	PostTime    time.Time        `json:"posttime"`
}

// FashionGallery -
type FashionGallery struct {
	//ID      bson.ObjectId  `bson:"_id,omitempty"`
	Stories string         `json:"stories"`
	Photos  []FashionPhoto `json:"photos"`
}

// FashionPhoto -
type FashionPhoto struct {
	Src string `json:"src"`
}
