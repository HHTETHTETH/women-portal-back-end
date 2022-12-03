package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Movie -
type Movie struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Board       string        `json:"board"`
	Trailer     string        `json:"trailer"`
	Release     string        `json:"release"`
	Imdb        int           `json:"imdb"`
	Actors      []Actor       `json:"actor"`
	PostTime    time.Time     `json:"posttime"`
	Category    string        `json:"category"`
}

// Actor -
type Actor struct {
	ActorName string `json:"actorname"`
	Photo     string `json:"photo"`
}
