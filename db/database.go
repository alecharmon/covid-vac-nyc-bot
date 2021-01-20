package db

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var singleton *gorm.DB
var once sync.Once

func GetDb() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		singleton = db
	})
	return singleton
}
