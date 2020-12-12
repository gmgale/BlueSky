﻿![Go logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

# BlueSky

BlueSky is an API that downloads an image depending on the current weather of a queried location.

The current weather endpoint is {host:port}/currentweather/{City}/{image Size}

Command line flags can be used:
* To set the host ("localhost" default). I.e -host 127.0.0.23. 
* To vary the API port ("9090" default). I.e -port 1234.

City should always be capitilized, the size should not.

Image size options are:
* original  
* large2x
* large
* medium
* small
* portrait
* landscape
* tiny

Example:

"http://localhost:9090/currentweather/Lisbon/large"

Will give the plain/text response: "The weather is Cloudy in Lisbon. Searching for images of Lisbon Cloudy.
Image pexels-photo-5959231.jpeg has been downloaded to the root folder.
Please credit the photographer Soulful Pizza / https://www.pexels.com/@soulful-pizza-2080276.

The image will then be saved to the same directory as the executable.
___


 Whenever you are using the service for your API, make sure to show a prominent link to Pexels. You can use a text link (e.g. "Photos provided by Pexels") or a link with their logo.

Always credit the photographers when possible (e.g. "Photo by John Doe on Pexels" with a link to the photo page on Pexels). 