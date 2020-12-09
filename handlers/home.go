package handlers

import (
	"fmt"
	"net/http"

	"github.com/gmgale/BlueSky/apikeys"
)

func Home(rw http.ResponseWriter, h *http.Request) {
	x := apikeys.LocalAPIKeys
	fmt.Printf("%s\n%s\n", x["weather"], x["images"])
}
