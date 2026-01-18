package handlers

import (
	"ponglockdown/server/abuse"
	"ponglockdown/server/storage"
)

var Store storage.LeaderboardStore = storage.NewMemoryStore()

var deviceBucket = abuse.NewTokenBucket(45, 25)
var ipBucket     = abuse.NewTokenBucket(60, 30)
