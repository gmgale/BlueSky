package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gmgale/BlueSky/apikeys"
	"github.com/gorilla/mux"
)

var GlobalWeatherResp weatherResp

func WeatherMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
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

		defer resp.Body.Close()

		text, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(rw, "Unable to read response.", http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal([]byte(text), &GlobalWeatherResp)
		if err != nil {
			http.Error(rw, "Error unmarshalling weather JSON.", http.StatusInternalServerError)
			return
		}

		//	Pretty print the weather to the console
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, text, "", "\t")
		if err != nil {
			log.Println("JSON parse error: ", err)
			return
		}
		log.Println("JSON response: ", string(prettyJSON.Bytes()))

		next.ServeHTTP(rw, r)
	})
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
