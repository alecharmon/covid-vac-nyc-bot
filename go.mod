module github.com/alecharmon/covid-vac-nyc-bot

// +heroku goVersion go1.15
go 1.15

require (
	github.com/dghubble/go-twitter v0.0.0-20201011215211-4b180d0cc78d
	github.com/dghubble/oauth1 v0.7.0
	github.com/joho/godotenv v1.3.0
	gorm.io/driver/postgres v1.0.6 // indirect
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.11
)
