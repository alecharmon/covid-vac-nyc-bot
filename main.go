package main

import (
	"fmt"

	"github.com/alecharmon/covid-vac-nyc-bot/sites"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	for _, v := range sites.GetSites() {
		fmt.Println(v.Avaliable())
	}

	// fmt.Println("Go-Twitter Bot v0.01")
	// creds := twitter.Credentials{
	// 	AccessToken:       os.Getenv("ACCESS_TOKEN"),
	// 	AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	// 	ConsumerKey:       os.Getenv("CONSUMER_KEY"),
	// 	ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	// }

	// fmt.Printf("%+v\n", creds)

	// client, err := twitter.getClient(&creds)
	// if err != nil {
	// 	log.Println("Error getting Twitter Client")
	// 	log.Println(err)
	// }

	// // Print out the pointer to our client
	// // for now so it doesn't throw errors
	// fmt.Printf("%+v\n", client)

}
