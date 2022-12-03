package utils

import (
	m "back/models"
	mgo "gopkg.in/mgo.v2"
	"time"
)

//ClickMapCollection is a mapping of Mongo collection of ClickMap
var ClickMapCollection *mgo.Collection

//InsertClickMap creates a record in collection ClickMap
func InsertClickMap(data m.ClickMap) error {
	/*t := time.Now()
	data.PostTimestamp = t.Unix()
	data.PostTime = t.Format(time.RFC3339)*/
	data.PostTime = time.Now()
	err := ClickMapCollection.Insert(&data)
	return err
}
