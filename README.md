﻿![Go logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

# BlueSky

BlueSky is an API that downloads an image depending on the current weather of a queried location.

Endpoints are localhost:9090/currentweather/{City}/{image Size}

Images are saved to the same directory as the executable.

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
will give the plain/text response: "The weather is Cloudy in Lisbon. Searching for images of Lisbon Cloudy.
Image pexels-photo-5959231.jpeg has been downloaded to the root folder.
Please credit the photographer Soulful Pizza / https://www.pexels.com/@soulful-pizza-2080276.
___


 Whenever you are using the service for your API, make sure to show a prominent link to Pexels. You can use a text link (e.g. "Photos provided by Pexels") or a link with their logo.

Always credit the photographers when possible (e.g. "Photo by John Doe on Pexels" with a link to the photo page on Pexels). 