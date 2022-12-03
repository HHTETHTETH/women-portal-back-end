package utils

import (
	m "back/models"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CookCollection -
var CookCollection *mgo.Collection

// FoodCoolection -
var FoodCoolection *mgo.Collection

// FindAllRecipes -
func FindAllRecipes() ([]m.Cook, error) {
	var datas []m.Cook
	err := CookCollection.Find(bson.M{}).Sort("-posttime").All(&datas)
	return datas, err
}

// FindRecipesByFood -
func FindRecipesByFood(food string) ([]m.Cook, error) {
	var datas []m.Cook
	err := CookCollection.Find(bson.M{"foodcategory": food}).Sort("-posttime").All(&datas)
	return datas, err
}

// InsertRecipes -
func InsertRecipes(data m.Cook) error {
	/*t := time.Now()
	data.PostTime = t.Unix()*/
	data.PostTime = time.Now()
	err := CookCollection.Insert(&data)
	return err
}

// DeleteRecipes - modified Remove to RemoveId (need ID only for deleting)
func DeleteRecipes(data m.Cook) error {
	err := CookCollection.RemoveId(&data.ID)
	return err
}

// UpdateRecipes -
func UpdateRecipes(data m.Cook) error {
	err := CookCollection.UpdateId(data.ID, &data)
	return err
}

// FindAllFoodCate -
func FindAllFoodCate() ([]m.FoodCategory, error) {
	var datas []m.FoodCategory
	//fmt.Println(" insdie FindAllFoodCate")
	err := FoodCoolection.Find(bson.M{}).All(&datas)
	//fmt.Println(" err ", err)
	return datas, err
}

// InsertFoodCate -
func InsertFoodCate(data m.FoodCategory) error {
	err := FoodCoolection.Insert(&data)
	return err
}

// DeleteFoodCate - modified Remove to RemoveId (need ID only for deleting)
func DeleteFoodCate(data m.FoodCategory) error {
	err := FoodCoolection.RemoveId(&data.ID)
	return err
}

// UpdateFoodCate -
func UpdateFoodCate(data m.FoodCategory) error {
	err := FoodCoolection.UpdateId(data.ID, &data)
	return err
}
