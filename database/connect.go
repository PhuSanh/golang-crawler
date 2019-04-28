package database

import (
	"github.com/globalsign/mgo"
	"go-crawler/config"
	"log"
)

var MongoDB *mgo.Database

func NewConn() (db *mgo.Database) {
	mongoConfig := config.Config.MongoDatabase

	session, err := mgo.Dial(mongoConfig.Host)
	if err != nil {
		log.Printf("Lỗi kết nối: %v\n", err)
		return
	}
	//defer session.Close()

	log.Println("Connected database go-crawler successful")

	db = session.DB(mongoConfig.DatabaseName)
	if db == nil {
		log.Printf("DB %s not found !!!\n", mongoConfig.DatabaseName)
		return
	}

	return

}
