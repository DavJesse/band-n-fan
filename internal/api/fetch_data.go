package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	Index     int      `json:"index"`
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	Index     int      `json:"index"`
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Relation struct {
	Index          int                 `json:"index"`
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Data struct {
	Artists   []Artist
	Locations []Location
	Dates     []Date
	Relations []Relation
}

var (
	artists   []Artist
	locations []Location
	dates     []Date
	relations []Relation
)

var apiURL = "https://groupietrackers.herokuapp.com/api"

func fetchData(url string, target interface{}) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to fetch data %s: %v", url, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read data: err")
	}
	return json.Unmarshal(body, &target)
}

func LoadData() (Data, error) {
	var err error
	if err = fetchData(apiURL+"/artists", &artists); err != nil {
		return Data{}, err
	}
	if err = fetchData(apiURL+"/locations", &locations); err != nil {
		return Data{}, err
	}
	if err = fetchData(apiURL+"/dates", &dates); err != nil {
		return Data{}, err
	}
	if err = fetchData(apiURL+"/relations", &relations); err != nil {
		return Data{}, err
	}

	data := Data{
		Artists: artists,
		Locations: locations,
		Dates: dates,
		Relations: relations,
	}
	return data, nil
}
