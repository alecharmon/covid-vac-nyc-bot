package sites

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	avaliable string = "AA"
)

type Sites struct {
	Sites []Site `json:"providerList"`
}

type Site struct {
	Name     string `json:"providerName"`
	Location string `json:"address"`
	Status   string `json:"availableAppointments"`
}

func (s Site) Avaliable() bool {
	return s.Status == avaliable
}

func GetSites() []Site {

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
