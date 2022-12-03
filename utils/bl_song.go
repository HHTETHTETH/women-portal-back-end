package utils

import (
	m "back/models"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SingerCollection -
var SingerCollection *mgo.Collection

// SongCollection -
var SongCollection *mgo.Collection

// InsertSinger -
func InsertSinger(data m.Singer) error {
	err := SingerCollection.Insert(&data)
	return err
}

// DeleteSinger - modified Remove to RemoveId (need ID only for deleting)
func DeleteSinger(data m.Singer) error {
	err := SingerCollection.RemoveId(&data.ID)
	return err
}

// UpdateSinger -
func UpdateSinger(data m.Singer) error {
	err := SingerCollection.UpdateId(data.ID, &data)
	return err
}

// FindAllSingers -
func FindAllSingers() ([]m.Singer, error) {
	var datas []m.Singer
	err := SingerCollection.Find(bson.M{}).Sort("name").All(&datas)
	return datas, err
}

// InsertSong -
func InsertSong(data m.Song) error {
	/*t := time.Now()
	data.PostTime = t.Unix()*/
	data.PostTime = time.Now()
	err := SongCollection.Insert(&data)
	return err
}

// DeleteSong - modified Remove to RemoveId (need ID only for deleting)
func DeleteSong(data m.Song) error {
	err := SongCollection.RemoveId(&data.ID)
	return err
}

// UpdateSong -
func UpdateSong(data m.Song) error {
	err := SongCollection.UpdateId(data.ID, &data)
	return err
}

// FindAllSongs -
func FindAllSongs() ([]m.Song, error) {
	var datas []m.Song
	err := SongCollection.Find(bson.M{}).Sort("-posttime").All(&datas)
	return datas, err
}

// FindSongsBySinger -
func FindSongsBySinger(name string) ([]m.Song, error) {
	var datas []m.Song
	err := SongCollection.Find(bson.M{"singer.name": name}).Sort("-posttime").All(&datas)
	return datas, err
}
