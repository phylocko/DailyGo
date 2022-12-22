package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func CreateConnection() {
	sqliteDB, err := gorm.Open(sqlite.Open("dailygo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect DB", err)
	}

	Db = sqliteDB
}
