package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gmgale/BlueSky/apikeys"
)

// CurrentWeather submits a GET request to the weather platform for data.
func GetImage(rw http.ResponseWriter, r *http.Request) {

	w := &GlobalWeatherResp
	APIkey := apikeys.LocalAPIKeys["images"]
	baseURL := "https://api.pexels.com/v1/search?query="
	perPage := "&per_page=1"
	URL := baseURL + w.Name + " " + w.Weather[0].Main + perPage

	log.Println("Making a request to: ", URL)
	log.Println("Using credentials: ", APIkey)

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Printf("%v", err)
		http.Error(rw, "Unable to fetch image data.", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", APIkey)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%v", err)
		http.Error(rw, "Unable to fetch image data.", http.StatusBadRequest)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(rw, "Unable to read response.", http.StatusInternalServerError)
		return
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	defer resp.Body.Close()

	var data imageData

	err = json.Unmarshal([]byte(bodyBytes), &data)
	if err != nil {
		log.Println("Error unmarshalling image JSON: ", err)
		http.Error(rw, "Error unmarshalling image JSON.", http.StatusInternalServerError)
		return
	}

	//	Pretty print the image data to the console
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, bodyBytes, "", "\t")
	if err != nil {
		log.Println("JSON parse error: ", err)
		return
	}
	log.Println("JSON response: ", string(prettyJSON.Bytes()))
}

type imageData struct {
	TotalResults int `json:"total_results"`
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	Photos       []struct {
		ID              int    `json:"id"`
		Width           int    `json:"width"`
		Height          int    `json:"height"`
		URL             string `json:"url"`
		Photographer    string `json:"photographer"`
		PhotographerURL string `json:"photographer_url"`
		PhotographerID  int    `json:"photographer_id"`
		Src             struct {
			Original  string `json:"original"`
			Large2x   string `json:"large2x"`
			Large     string `json:"large"`
			Medium    string `json:"medium"`
			Small     string `json:"small"`
			Portrait  string `json:"portrait"`
			Landscape string `json:"landscape"`
			Tiny      string `json:"tiny"`
		} `json:"src"`
		Liked bool `json:"liked"`
	} `json:"photos"`
	NextPage string `json:"next_page"`
}
