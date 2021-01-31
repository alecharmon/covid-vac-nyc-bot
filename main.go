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

	availableSites := []*sites.Site{}
	unavailableSites := []*sites.Site{}
	db.GetDb().AutoMigrate(&sites.Site{})
	for _, v := range sites.GetSites() {
		if v.IsNewSite() || sites.GetFromName(v.Name).Avaliable() != v.Avaliable() {
			if v.Avaliable() {
				availableSites = append(availableSites, v)
			} else {
				unavailableSites = append(unavailableSites, v)
			}
		}
	}

	// Available
	if len(availableSites) > 0 {
		messages := []string{}
		message := "The following sites now have appointments available: \n "
		for _, s := range availableSites {
			newSite := s.IsNewSite()
			prefix := ""
			if newSite {
				prefix = "*New Site* "
			}
			message += prefix + s.ToString() + "\n "
			if newSite {
				db.GetDb().Create(s)
			} else {
				record := sites.GetFromName(s.Name)
				if record == nil {
					log.Println("Not a new site but not retrivable by name " + s.ToString())
				}
				record.Status = s.Status
				db.GetDb().Save(&record)
			}
			if len(message) > 280 {
				messages = append(messages, message)
				message = ""
			}
		}
		tail := "💉💊🎊 for more info https://am-i-eligible.covid19vaccine.health.ny.gov/"
		if len(message)+len(tail) >= 280 {
			messages = append(messages, message)
			message = ""
		}
		message += tail
		messages = append(messages, message)
		config := twitter.NewParams()
		for _, message := range messages {
			tweet, resp, err := client.Statuses.Update(message, config)
			if err != nil {
				log.Fatal(message, err, resp)
			}
			config.InReplyToStatusID = tweet.ID
		}
	}
}

func dropTableIfExists(db *gorm.DB, dst interface{}) {
	if db.Migrator().HasTable(dst) {
		db.Migrator().DropTable(dst)
	}
}
