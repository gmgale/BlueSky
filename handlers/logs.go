package handlers

import (
	"fmt"
	"net/http"

	"github.com/gmgale/BlueSky/ratelimit"
)

// Logs prints the current memory store of rate limiting logs to the browser.
func Logs(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "%v \n", ratelimit.UserLog)

}
