package models

import "gopkg.in/mgo.v2/bson"

type Admin struct {
	ID          	bson.ObjectId `bson:"_id,omitempty"`
	UserName 		string `json:"userName"`
	Password 		string `json:"password"`
}
