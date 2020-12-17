﻿![Go logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

# <u>BlueSky</u>

BlueSky is an API service that downloads professionally taken images of a location, depending on the locations current weather.

This is achieved by leveraging the [Pexels](https://www.pexels.com/) and [OpenWeatherMap](https://openweathermap.org/) API services. You will need to provide your own API keys for these services as decribed below.
___

<u>Endpoints</u>

The current weather endpoint is: 
```
{host:port}/currentweather/{City}/{image size}
```
The City name should always be capitilized, the size should not.

Image size options are:
* original  
* large2x
* large
* medium
* small
* portrait
* landscape
* tiny

Example request:
```
127.0.0.24:9090/currentweather/Lisbon/large
```

Will give the plain/text response: "The weather is Cloudy in Lisbon. Searching for images of Lisbon Cloudy.
Image pexels-photo-5959231.jpeg has been downloaded to the photos folder.
Please credit the photographer Soulful Pizza / https://www.pexels.com/@soulful-pizza-2080276.

The image will then be saved into the */photos* directory (created on boot and deleted on shutdown).
Please note images can be up to ~10Mb.

The logs endpoint is:
```
{host:port}/logs 
```
This will display the current memory log of the rate-limiting middleware if enabled.
___

<u>Command line flags</u>

The optional flags can be used:
* *-host*: To set the host ("localhost" default). 
* *-port*: To vary the API port ("9090" default).
* *-limit:* To enable rate limiting ("-1" / off default).
* *-test*: To enable test mode ("false" default).
     Test mode disables calls to external APIs to avoid superfluous requests during development.

Powershell example:
```
.\main.exe -host 127.0.0.24 -port 8080 -limit 200
```
___

<u>API Keys</u>

You will need to provide your own API keys for the [Pexels](https://www.pexels.com/) and [OpenWeatherMap](https://openweathermap.org/) API services. These should then be accessable from a map as shown:
```
var LocalAPIKeys = map[string]string{
	"weather": "1a2b3c4d",
	"images":  "5e6f7g8h",
}
```
___

<u>Crediting Photographers</u>

This API is not affiliated with Pexels.

 Whenever you are using the service for your API, make sure to show a prominent link to Pexels. You can use a text link (e.g. "Photos provided by Pexels") or a link with their logo.

Always credit the photographers when possible (e.g. "Photo by John Doe on Pexels" with a link to the photo page on Pexels). 