package utils

import (
	m "back/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TansCollection -  Empolyee CRUD operation
var TansCollection *mgo.Collection

// FindAllCBTrans - All Avaliable User
func FindAllCBTrans() ([]m.CallBack, error) {
	var datas []m.CallBack
	err := TansCollection.Find(bson.M{}).All(&datas)
	return datas, err
}

// FindCBTansByTransID -
func FindCBTansByTransID(trans string) (m.CallBack, error) {
	var data m.CallBack
	err := TansCollection.Find(bson.M{"transID": trans}).One(&data)
	return data, err
}

// InsertCBTran - Create new Unit
func InsertCBTran(data m.CallBack) error {
	err := TansCollection.Insert(&data)
	return err
}

// DeleteCBTran - Delete Unit
func DeleteCBTran(data m.CallBack) error {
	err := TansCollection.Remove(&data)
	return err
}

// UpdateCBTran -
func UpdateCBTran(data m.CallBack) error {
	err := TansCollection.UpdateId(data.ID, &data)
	return err
}
