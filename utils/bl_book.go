package utils

import (
	m "back/models"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BookCollection -
var BookCollection *mgo.Collection

// FindAllBooks -
func FindAllBooks() ([]m.Book, error) {
	var datas []m.Book
	err := BookCollection.Find(bson.M{}).Sort("-posttime").All(&datas)
	return datas, err
}

// InsertBook -
func InsertBook(data m.Book) error {
	/*t := time.Now()
	data.PostTime = t.Unix()*/
	data.PostTime = time.Now()
	err := BookCollection.Insert(&data)
	return err
}

// UpdateBook - 
func UpdateBook(data m.Book) error {
	err := BookCollection.UpdateId(&data.ID, &data)
	return err
}

//DeleteBook - modified Remove to RemoveId (need ID only for deleting)
func DeleteBook(data m.Book) error {
	err := BookCollection.RemoveId(&data.ID)
	return err
}
