package ratelimit

import "time"

// GlobalRateLimit should e set with the command line flag -limit.
var GlobalRateLimit int

// UserLog is used to keep track of Address: Requests remaining, Time.
var UserLog UserMap

type User struct {
	ReqRemain int
	LastReq   time.Time
}

// UserMap is used for efficient record of
// User address: Requests remaining, Time.
type UserMap map[string]User

// SessionStorage allocates memory on the stack
// for user logs.
func SessionStorage() *UserMap {
	store := make(UserMap, 1000)
	storePtr := &store
	return storePtr
}
