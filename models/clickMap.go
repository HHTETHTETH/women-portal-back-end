package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//CallBack describes a model on Mongo DB
// It contains a log of vendor`s click id mapped to msisdn and MA`s transaction id
type ClickMap struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	Vendor        string        `json:"vendor"`
	Click         string        `json:"clickId"`
	SequenceNo    string        `json:"sequenceNo"`
	CallingParty  string        `json:"callingParty"`
	ServiceID     string        `json:"serviceId"`
	PostTimestamp int64         `json:"postTimestamp,omitempty"`
	PostTime      time.Time     `json:"postTime,omitempty"`
}
