﻿![Go logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

# BlueSky

BlueSky is an API that downloads an image depending on the weather of a queried location.

Endpoints are localhost:9090/currentweather/{City}/{Image Size}

Images are saved to the same directory as the executable.

Example:
    http://localhost:9090/currentweather/Lisbon/large

City should be capitilized.

Image size options are:
* original  
* large2x
* large
* medium
* small
* portrait
* landscape
* tiny
