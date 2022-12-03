package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Article -
type Article struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Title       string        `json:"title"`
	Date        string        `json:"date"`
	Categories  []Categories  `json:"categories"`
	Image       string        `json:"image"`
	Description string        `json:"description"`
	Author      string        `json:"author"`
	Count       Count         `json:"count"`
	Contents    []Contents    `json:"contents"`
	PostTime    time.Time     `json:"posttime"`
}

// Categories -
type Categories struct {
	Label string `bson:"label"`
}

// Count -
type Count struct {
	Like  int `bson:"like"`
	Count int `bson:"count"`
}

// Contents -z
type Contents struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Images      []ContentImages `json:"images"`
}

// ContentImages -
type ContentImages struct {
	Src   string `json:"src"`
	Title string `json:"title"`
}
