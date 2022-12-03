package utils

import (
	m "back/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserCollection -  Empolyee CRUD operation
var UserCollection *mgo.Collection

// FindAllUsers - All Avaliable User
func FindAllUsers() ([]m.User, error) {
	var datas []m.User
	err := UserCollection.Find(bson.M{}).All(&datas)
	return datas, err
}

// FindUserByPhone -
func FindUserByPhone(phone string) (m.User, error) {
	var data m.User
	err := UserCollection.Find(bson.M{"phone": phone}).One(&data)
	return data, err
}

// InsertUser - Create new Unit
func InsertUser(data m.User) error {
	err := UserCollection.Insert(&data)
	return err
}

// DeleteUser - Delete Unit
func DeleteUser(data m.User) error {
	err := UserCollection.Remove(&data)
	return err
}

// UpdateUser -
func UpdateUser(data m.User) error {
	err := UserCollection.UpdateId(data.ID, &data)
	return err
}
