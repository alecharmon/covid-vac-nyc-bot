package main

import (
	"log"
	"os"

	"github.com/alecharmon/covid-vac-nyc-bot/db"
	"github.com/alecharmon/covid-vac-nyc-bot/sites"
	"github.com/alecharmon/covid-vac-nyc-bot/twitter"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	creds := twitter.Credentials{
		AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("TWITTER_CONSUMER_SECRET"),
	}

	client, err := twitter.GetClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	db.GetDb().AutoMigrate(&sites.Site{})
	for _, v := range sites.GetSites() {
		log.Println("Site ", v)
		lastRecord := sites.GetFromName(v.Name)
		if lastRecord.ID == "" {
			log.Println("New Site ", v)
			v.ID = sites.GetKey(v.Name)
			db.GetDb().Create(&v)
			client.Statuses.Update("New Site: "+v.ToString(), nil)
		} else if v.Avaliable() && lastRecord.Avaliable() == false {
			log.Println("New Availability @ Site ", lastRecord)
			lastRecord.Status = v.Status
			db.GetDb().Save(&lastRecord)
			client.Statuses.Update("Site Update: "+v.ToString(), nil)
		}

	}
}

func dropTableIfExists(db *gorm.DB, dst interface{}) {
	if db.Migrator().HasTable(dst) {
		db.Migrator().DropTable(dst)
	}
}
