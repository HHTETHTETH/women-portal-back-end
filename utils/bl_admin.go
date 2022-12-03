package utils

import (
	m "back/models"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//AdminCollection is a mapping of Mongo collection of ClickMap
var AdminCollection *mgo.Collection

// FindAllAdmins - All Available Admin
func FindAllAdmins() ([]m.Admin, error) {
	var admins []m.Admin
	err := AdminCollection.Find(bson.M{}).All(&admins)
	return admins, err
}

// FindAdminByAdminInfo - get only one admin
func FindAdminByInfo(userName, password string) (m.Admin, error) {
	var admin m.Admin
	err := AdminCollection.Find(bson.M{"username": userName, "password": password}).One(&admin)
	//fmt.Println("err : ", err, "Admin : ", admin)
	return admin, err
}

func FindAdminByPwd(password string) (m.Admin, error) {
	var admin m.Admin
	err := AdminCollection.Find(bson.M{"password": password}).One(&admin)
	//fmt.Println("err : ", err, "Admin : ", admin)
	return admin, err
}

// InsertAdmin - Create new Unit
func InsertAdmin(data m.Admin) error {
	var flag bool
	admins, err := FindAllAdmins()
	if err != nil{
		return err
	}
	for _, admin := range admins{
		if admin.Password == data.Password{
			flag = true
			break
		}
	}
	if flag == false{
		err := AdminCollection.Insert(&data)
		return err
	}else {
		return errors.New("This Password is already used by other user!")
	}
}

// DeleteAdmin - Delete Unit
func DeleteAdmin(data m.Admin) error {
	err := AdminCollection.Remove(&data)
	return err
}

// UpdateAdmin - Update Unit
func UpdateAdmin(data m.Admin) error {
	err := AdminCollection.UpdateId(data.ID, &data)
	fmt.Println("err : ", err, "data: ", data)
	return err
}

