package utils

import (
	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

const (
	DTATBASE = "khintkabardb02"
)

var DB *mgo.Database

// InitDriver - Init the MongoDb Connection
func InitDriver(mongoURL string) error {
	//mongodb://mongo:27018
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		panic(err)
	}
	DB = session.DB(DTATBASE)
	UserCollection = DB.C("user")
	TansCollection = DB.C("callbacks")
	ArticleCollection = DB.C("article")
	PromotionCollection = DB.C("promos")
	CookCollection = DB.C("cook")
	FoodCoolection = DB.C("food")
	FashionCollection = DB.C("fashion")
	SingerCollection = DB.C("singer")
	SongCollection = DB.C("song")
	BookCollection = DB.C("book")
	MovieCollection = DB.C("movies")
	InspireCollection = DB.C("inspire")
	DIYCollection = DB.C("diy")
	ClickMapCollection = DB.C("clickMap")
	AdminCollection = DB.C("admin")
	return nil
}
