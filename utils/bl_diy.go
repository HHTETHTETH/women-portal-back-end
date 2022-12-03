package utils

import (
	m "back/models"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DIYCollection -
var DIYCollection *mgo.Collection

// FindAllDIYS - All Avaliable User
func FindAllDIYS() ([]m.Diy, error) {
	var datas []m.Diy
	err := DIYCollection.Find(bson.M{}).Sort("-posttime").All(&datas)
	return datas, err
}

// InsertDIY -
func InsertDIY(data m.Diy) error {
	/*t := time.Now()
	data.PostTime = t.Unix()*/
	data.PostTime = time.Now()
	err := DIYCollection.Insert(&data)
	return err
}

// DeleteDIY - modified Remove to RemoveId (need ID only for deleting)
func DeleteDIY(data m.Diy) error {
	err := DIYCollection.RemoveId(&data.ID)
	return err
}

// UpdateDIY -
func UpdateDIY(data m.Diy) error {
	err := DIYCollection.UpdateId(data.ID, &data)
	return err
}
