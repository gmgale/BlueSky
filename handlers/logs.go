package handlers

import (
	"fmt"
	"net/http"

	"github.com/gmgale/BlueSky/ratelimit"
)

func Logs(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(rw, "%v ", ratelimit.UserLog)

}
