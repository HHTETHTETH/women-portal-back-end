package utils

import (
	m "back/models"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// FashionCollection -
var FashionCollection *mgo.Collection

// FindAllFashions -
func FindAllFashions() ([]m.Fashion, error) {
	var datas []m.Fashion
	err := FashionCollection.Find(bson.M{}).Sort("-posttime").All(&datas)
	return datas, err
}

// InsertFashion -
func InsertFashion(data m.Fashion) error {
	/*t := time.Now()
	data.PostTime = t.Unix()*/
	data.PostTime = time.Now()
	err := FashionCollection.Insert(&data)
	return err
}

// DeleteFashion - modified Remove to RemoveId (need ID only for deleting)
func DeleteFashion(data m.Fashion) error {
	err := FashionCollection.RemoveId(&data.ID)
	return err
}

// UpdateFashion -
func UpdateFashion(data m.Fashion) error {
	err := FashionCollection.UpdateId(data.ID, &data)
	return err
}
