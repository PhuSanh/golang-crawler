package models

import (
	"github.com/globalsign/mgo/bson"
	"go-crawler/database"
	"log"
)

type News struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Link string `json:"link" bson:"link"`
	Thumb string `json:"thumb" bson:"thumb"`
	Description string `json:"description" bson:"description"`
	Content string `json:"content" bson:"content"`
}

func Insert(news *News) (err error) {
	collection := database.MongoDB.C("news")
	err = collection.Insert(news)
	if err != nil {
		log.Println("[Error]: Insert news", err.Error())
		return
	}
	return
}

func GetList() (listNews []News, err error) {
	collection := database.MongoDB.C("news")

	err = collection.Find(nil).All(&listNews)
	if err != nil {
		log.Println("[Error]: Get list news", err.Error())
		return listNews, err
	}

	return listNews, nil
}

func UpdateById(news News) (err error) {
	collection := database.MongoDB.C("news")
	err = collection.UpdateId(news.ID, news)

	if err != nil {
		log.Println("[Error]: Update news", err.Error())
		return
	}

	return
}