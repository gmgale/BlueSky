package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gmgale/BlueSky/apikeys"
	download "github.com/gmgale/BlueSky/data"
	"github.com/gorilla/mux"
)

// CurrentWeather submits a GET request to the weather platform for data.
func GetImage(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	imgSize := vars["imgSize"]

	w := GlobalWeatherResp
	APIkey := apikeys.LocalAPIKeys["images"]
	baseURL := "https://api.pexels.com/v1/search?query="
	perPage := "&per_page=1"
	URL := baseURL + w.Name + "-" + w.Weather[0].Main + perPage

	fmt.Fprintf(rw, "Searching for images of %s %s.\n", w.Name, w.Weather[0].Main)
	log.Println("Making a request to: ", URL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Printf("%v", err)
		http.Error(rw, "Unable to fetch image data.", http.StatusInternalServerError)
		return
	}

	req.Header.Add("Authorization", APIkey)
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
	defer resp.Body.Close()

	var data imageData
	err = json.Unmarshal([]byte(bodyBytes), &data)
	if err != nil {
		http.Error(rw, "Error unmarshalling image JSON.", http.StatusInternalServerError)
		return
	}

	// This map is used to select the correct URL depending on the
	// endpoint varible imgSize
	paths := map[string]string{
		"original":  data.Photos[0].Src.Original,
		"large2x":   data.Photos[0].Src.Large2x,
		"large":     data.Photos[0].Src.Large,
		"medium":    data.Photos[0].Src.Medium,
		"small":     data.Photos[0].Src.Small,
		"portrait":  data.Photos[0].Src.Portrait,
		"landscape": data.Photos[0].Src.Landscape,
		"tiny":      data.Photos[0].Src.Tiny,
	}

	// Extract URL from json
	fileURL := paths[imgSize]

	fname := ""
	// Download the image to the root folder
	fname, err = download.DownloadFile(fileURL)
	if err != nil {
		log.Println("Error downloding image.", err)
		http.Error(rw, "Error downloding image.", http.StatusInternalServerError)
	}

	fmt.Fprintf(rw, "Image %s has been downloaded to the root folder.\n", fname)
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
