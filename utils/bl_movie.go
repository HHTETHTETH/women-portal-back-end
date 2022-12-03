package utils

import (
	m "back/models"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MovieCollection -
var MovieCollection *mgo.Collection

// FindAllMovies -
func FindAllMovies() ([]m.Movie, error) {
	var datas []m.Movie
	err := MovieCollection.Find(bson.M{}).Sort("-posttime").All(&datas)
	return datas, err
}

// InsertMovie -
func InsertMovie(data m.Movie) error {
	/*t := time.Now()
	data.PostTime = t.Unix()*/
	data.PostTime = time.Now()
	err := MovieCollection.Insert(&data)
	return err
}

// DeleteMovie - modified Remove to RemoveId (need ID only for deleting)
func DeleteMovie(data m.Movie) error {
	err := MovieCollection.RemoveId(&data.ID)
	return err
}

// UpdateMovie -
func UpdateMovie(data m.Movie) error {
	err := MovieCollection.UpdateId(data.ID, &data)
	return err
}
