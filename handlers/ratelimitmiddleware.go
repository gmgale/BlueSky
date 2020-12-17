package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gmgale/BlueSky/ratelimit"
)

// RatelimitMiddleware is hit first when the endpoint is requsted.
// It returns immediately if the command line flag of -rate is -1/off.
// If enabled, a log is kept on the stack of user interactions and
// limiting is applied accordingly.
func RatelimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		if ratelimit.GlobalRateLimit == -1 {
			// Rate limit is off
			next.ServeHTTP(rw, r)
			return
		}

		t := time.Now()
		a := strings.Split(r.RemoteAddr, ":")
		addr := a[0]

		if val, ok := ratelimit.UserLog[addr]; ok {
			fmt.Printf("User %s is already in log.\n", addr)
			rr := val.ReqRemain
			lr := val.LastReq
			nt := t.Add(-time.Minute)

			// If time > 1 min, reset the log
			if nt.Sub(lr) > 1 {
				ratelimit.UserLog[addr] = ratelimit.User{
					ReqRemain: ratelimit.GlobalRateLimit - 1,
					LastReq:   time.Now(),
				}
				fmt.Fprintf(rw, "You have %d requests remaining.\n", ratelimit.GlobalRateLimit-1)
				next.ServeHTTP(rw, r)
				return
			}

			// If time < 1 min, check last log time and requests remaining
			if nt.Sub(lr) < 1 {
				// If remaining requests > 0 continue and add request to the log
				if rr > 0 {
					ratelimit.UserLog[addr] = ratelimit.User{
						ReqRemain: rr - 1,
						LastReq:   time.Now(),
					}

					fmt.Printf("You have %d requests remaining.\n", rr-1)
					fmt.Fprintf(rw, "You have %d requests remaining.\n", rr-1)
					next.ServeHTTP(rw, r)
					return
				}
				// if remaining requests == 0, display error
				if rr == 0 {
					fmt.Fprintf(rw, "Rate limit %d exceeded. Please wait 1 minute.\n", ratelimit.GlobalRateLimit)
					fmt.Printf("Rate limit %d exceeded. Please wait one minute.\n", ratelimit.GlobalRateLimit)
					next.ServeHTTP(rw, r)
					return
				}
			}

		} else {
			fmt.Printf("User %s is not in the log. Creating new entry.\n", addr)

			//create file
			newUser := ratelimit.User{
				ReqRemain: ratelimit.GlobalRateLimit - 1,
				LastReq:   time.Now(),
			}

			fmt.Printf("You have %d requests remaining.\n", ratelimit.GlobalRateLimit-1)
			fmt.Fprintf(rw, "You have %d requests remaining.\n", ratelimit.GlobalRateLimit-1)
			ratelimit.UserLog[addr] = newUser
		}

		next.ServeHTTP(rw, r)
	})
}
