package utils

import (
	m "back/models"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PromotionCollection -
var PromotionCollection *mgo.Collection

// FindAllPromotions - Skip() and Limit()
func FindAllPromotions() ([]m.Promotion, error) {
	var datas []m.Promotion
	err := PromotionCollection.Find(bson.M{}).Sort("-posttime").All(&datas)
	return datas, err
}

// FindPromotsbyDate -
func FindPromotsbyDate(date string) ([]m.Promotion, error) {
	var datas []m.Promotion
	err := PromotionCollection.Find(bson.M{"startdate": date}).All(&datas)
	//fmt.Println(" error (inside)", err)
	//	err := UserCollection.Find(bson.M{"phone": phone}).One(&data)
	return datas, err
}

// InsertPromotion -
func InsertPromotion(data m.Promotion) error {
	/*t := time.Now()
	data.PostTime = t.Unix()*/
	data.PostTime = time.Now()
	err := PromotionCollection.Insert(&data)
	return err
}

// implemented delete function
// DeletePromotion - modified Remove to RemoveId (need ID only for deleting)
func DeletePromotion(data m.Promotion) error {
	err := PromotionCollection.RemoveId(&data.ID)
	return err
}

// implemented update function
// UpdatePromotion -
func UpdatePromotion(data m.Promotion) error {
	err := PromotionCollection.UpdateId(data.ID, &data)
	return err
}
