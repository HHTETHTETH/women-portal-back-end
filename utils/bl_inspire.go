package utils

import (
	m "back/models"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// InspireCollection -
var InspireCollection *mgo.Collection

// FindAllInspires -
func FindAllInspires() ([]m.Inspire, error) {
	var datas []m.Inspire
	err := InspireCollection.Find(bson.M{}).Sort("-posttime").All(&datas)
	return datas, err
}

// InsertInspire -
func InsertInspire(data m.Inspire) error {
	/*t := time.Now()
	data.PostTime = t.Unix()*/
	data.PostTime = time.Now()
	err := InspireCollection.Insert(&data)
	return err
}

// DeleteInspire - modified Remove to RemoveId (need ID only for deleting)
func DeleteInspire(data m.Inspire) error {
	err := InspireCollection.RemoveId(&data.ID)
	return err
}

// UpdateInspire -
func UpdateInspire(data m.Inspire) error {
	err := InspireCollection.UpdateId(data.ID, &data)
	return err
}
