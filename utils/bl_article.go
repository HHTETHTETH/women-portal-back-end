package utils

import (
	m "back/models"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ArticleCollection -
var ArticleCollection *mgo.Collection

// FindAllArticles - All Avaliable User
func FindAllArticles() ([]m.Article, error) {
	var datas []m.Article
	err := ArticleCollection.Find(bson.M{}).Sort("-posttime").All(&datas)
	return datas, err
}

// InsertArticle -
func InsertArticle(data m.Article) error {
	/*t := time.Now()
	data.PostTime = t.Unix()*/
	data.PostTime = time.Now()
	err := ArticleCollection.Insert(&data)
	return err
}

// DeleteArticle - modified Remove to RemoveId (need ID only for deleting)
func DeleteArticle(data m.Article) error {
	err := ArticleCollection.RemoveId(&data.ID)
	return err
}

// UpdateArticle -
func UpdateArticle(data m.Article) error {
	err := ArticleCollection.UpdateId(data.ID, &data)
	return err
}
