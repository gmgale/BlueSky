package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gmgale/BlueSky/apikeys"
	"github.com/gorilla/mux"
)

// CurrentWeather submits a GET request to the weather platform for data.
func CurrentWeather(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]

	APIkey := apikeys.LocalAPIKeys["weather"]
	URL := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + APIkey

	resp, err := http.Get(URL)
	if err != nil {
		log.Printf("%v", err)
		http.Error(rw, "Unable to fetch weather data.", http.StatusBadRequest)
		return
	}

	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(rw, "Unable to read response.", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	res := &weatherResp{}
	err = json.Unmarshal([]byte(text), res)
	if err != nil {
		http.Error(rw, "Error unmarshalling JSON.", http.StatusInternalServerError)
		return
	}

	fmt.Printf("%v\n", res)
}

type results struct {
	responces []weatherResp
}

type weatherResp struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"discription"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp:max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}
