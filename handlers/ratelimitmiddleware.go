package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gmgale/BlueSky/ratelimit"
)

func RatelimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		if ratelimit.GlobalRateLimit == "-1" {
			// Limit is off
			fmt.Printf("WARNING: Rate-limiting is switched off.\n")
			fmt.Printf("Use commang line flag '-limit' to set.\n")
			return
		}
		
		wd, _ := os.Getwd()
		if wd != "data" {
			err := os.Chdir("/data")
			if err != nil {
				// fmt.Printf("RateLimitMiddleware: Error switching to data directory.\n%v\n", err)
				// fmt.Printf("Warning: Rate limiting may be disabled.\n")
			}
		}

		a := strings.Split(r.RemoteAddr, ":")
		addr := a[0]
		uFile := addr + "_log.tmp"

		newUser := ratelimit.User{
			RemAddress: addr,
			ReqRemain:  ratelimit.GlobalRateLimit,
		}

		if fileExists(uFile) {
			newUser.CheckLog()
			//open the file
			//check last log time and reqRemain
			//if time > 1 min ago; clear all log
			//if time < min ago:
			//	if rate < limit:
			//		write new log and close file
			//	if rate = limit:
			// display wait message
		} else {
			//create file

			newUser.MakeNewLog()
			//write log and close file
			//continue
		}
		next.ServeHTTP(rw, r)
	})
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
