package ratelimit

import "time"

var GlobalRateLimit int
var UserLog UserMap

type User struct {
	ReqRemain int
	LastReq   time.Time
}

type UserMap map[string]User

func SessionStorage() *UserMap {
	store := make(UserMap, 1000)
	storePtr := &store
	return storePtr
}
