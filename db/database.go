package db

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var singleton *gorm.DB
var once sync.Once

func GetDb() *gorm.DB {
	once.Do(func() {
		dsn := "host=" + os.Getenv("HOST")
		dsn += " user=" + os.Getenv("USER")
		dsn += " pasword=" + os.Getenv("PASWORD")
		dsn += " dbname=" + os.Getenv("DBNAME")
		dsn += " port=" + os.Getenv("PORT")

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
		}
		singleton = db
	})
	return singleton
}
