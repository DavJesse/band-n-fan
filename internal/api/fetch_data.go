package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relation     map[string][]string
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type LocationData struct {
	Index []Location `json:"index"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DateData struct {
	Index []Date `json:"index"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationData struct {
	Index []Relation `json:"index"`
}

type Data struct {
	Artists   []Artist
	Locations LocationData
	Dates     DateData
	Relations RelationData
}

var (
	artists   []Artist
	locations LocationData
	dates     DateData
	relations RelationData
)

var apiURL = "https://groupietrackers.herokuapp.com/api"

// fetchData is used to retrieve json data from a specified url...
// ... and storing data in go data structures.
func fetchData(url string, target interface{}) error {

	// Retrieve http response from url containing json data
	// Handle any errors encountered
	// Close response body at the termination of function
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from %s: %v", url, err)
	}
	defer resp.Body.Close()

	// Confirm http status code 200, to ensure everything is in order
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data %s: %v", url, resp.Status)
	}

	// Read response body, storing it in 'body'
	// Handle any errors encountered
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read data from %s: %v", url, err)
	}

	// Unmarshal json data in body, storing it in target
	return json.Unmarshal(body, &target)
}

// LoadData leverages fetch data to populate Data object with data from api
func LoadData() (Data, error) {
	var err error

	// Load data to 'ariststs', 'locations', 'dates', and 'relations'
	if err = fetchData(apiURL+"/artists", &artists); err != nil {
		return Data{}, err
	}
	if err = fetchData(apiURL+"/locations", &locations); err != nil {
		return Data{}, err
	}
	if err = fetchData(apiURL+"/dates", &dates); err != nil {
		return Data{}, err
	}
	if err = fetchData(apiURL+"/relation", &relations); err != nil {
		return Data{}, err
	}

	// consolidate the loaded data in 'data' for easier retrieval in the future
	data := Data{
		Artists:   artists,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}
	return data, nil
}
