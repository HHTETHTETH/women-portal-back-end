package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Cook -
type Cook struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Title        string        `json:"title"`
	MainPhoto    string        `json:"mainphoto"`
	MainVideo    string        `json:"mainvideo"`
	Author       string        `json:"author"`
	Date         string        `json:"date"`
	Credit       string        `json:"credit"`
	SlidePhotos  []SlidePhoto  `json:"slidephotos"`
	Description  string        `json:"description"`
	CookIngres   []CookIngre   `json:"cookingres"`
	CookSteps    []CookStep    `json:"cooksteps"`
	FoodCategory string        `json:"foodcategory"`
	PostTime     time.Time     `json:"posttime"`
}

// SlidePhoto -
type SlidePhoto struct {
	Src string `json:"src"`
}

// CookIngre -
type CookIngre struct {
	Icon      string `json:"icon"`
	Weight    string `json:"weight"`
	Unit      string `json:"unit"`
	IngreName string `json:"ingrename"`
}

// CookStep -
type CookStep struct {
	Title   string `json:"title"`
	Step    string `json:"step"`
	Process string `json:"process"`
}

// FoodCategory  -
type FoodCategory struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `json:"name"`
	IconPhoto string        `json:"iconphoto"`
}
