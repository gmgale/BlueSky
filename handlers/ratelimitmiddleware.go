package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gmgale/BlueSky/ratelimit"
)

func RatelimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		if ratelimit.GlobalRateLimit == -1 {
			// Rate limit is off
			next.ServeHTTP(rw, r)
		}

		t := time.Now()
		a := strings.Split(r.RemoteAddr, ":")
		addr := a[0]

		fmt.Printf("1. addr: %s\n", addr)
		fmt.Println("Logs:", ratelimit.UserLog)

		if val, ok := ratelimit.UserLog[addr]; ok {
			fmt.Printf("User %s is already in log.\n", addr)
			rr := val.ReqRemain
			lr := val.LastReq
			nt := t.Add(-time.Minute)

			// If its been > than 1 minute, reset the log
			if nt.Sub(lr) > 1 {
				fmt.Print("> than 1 minute since last call, resetting log.\n", addr)

				ratelimit.UserLog["addr"] = ratelimit.User{
					ReqRemain: ratelimit.GlobalRateLimit,
					LastReq:   time.Now(),
				}
				next.ServeHTTP(rw, r)
			}

			// If time < min ago, check last log time and requests remaining
			if nt.Sub(lr) < 1 {
				if rr == 0 {
					fmt.Fprintf(rw, "Rate limit %d exceeded. Please wait 1 minute.\n", ratelimit.GlobalRateLimit)
					fmt.Printf("Rate limit %d exceeded. Please wait one minute.\n", ratelimit.GlobalRateLimit)
					next.ServeHTTP(rw, r)
				}
				if rr > 0 {
					fmt.Printf("You have %d calls remaining.\n", rr)
					ratelimit.UserLog["addr"] = ratelimit.User{
						ReqRemain: rr - 1,
						LastReq:   time.Now(),
					}
					next.ServeHTTP(rw, r)
				}
			}
		} else {
			fmt.Printf("User %s is not in the log. Creating new entry.\n", addr)

			//create file
			newUser := ratelimit.User{
				ReqRemain: ratelimit.GlobalRateLimit,
				LastReq:   time.Now(),
			}
			ratelimit.UserLog[addr] = newUser
		}

		next.ServeHTTP(rw, r)
	})
}
