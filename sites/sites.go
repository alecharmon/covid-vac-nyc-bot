package sites

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/alecharmon/covid-vac-nyc-bot/db"
	"gorm.io/gorm"
)

const (
	avaliable string = "AA"
)

type Sites struct {
	Sites []*Site `json:"providerList"`
}

type Site struct {
	ID       string `gorm:"primaryKey"`
	Name     string ` json:"providerName"`
	Location string `json:"address"`
	Status   string `json:"availableAppointments"`
}

func GetKey(str string) string {
	return strings.Join(strings.Fields(str), "")
}
func (s Site) Avaliable() bool {
	return s.Status == avaliable
}

func (s *Site) ToString() string {
	avaliable := "does not have availability for vaccine appointments"
	if s.Avaliable() {
		avaliable = "has availability!💊🎊 more info https://am-i-eligible.covid19vaccine.health.ny.gov/"
	}
	return fmt.Sprintf("%s @ %s %s", s.Name, s.Location, avaliable)
}

func GetFromName(name string) *Site {
	site := &Site{}
	db.GetDb().Where("id = ? ", GetKey(name)).First(&site)
	return site
}

func (s *Site) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = GetKey(s.Name)

	return
}

func GetSites() []*Site {
	url := "https://am-i-eligible.covid19vaccine.health.ny.gov/api/list-providers"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	sites := Sites{}
	jsonErr := json.Unmarshal(body, &sites)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(sites)

	return sites.Sites
}
