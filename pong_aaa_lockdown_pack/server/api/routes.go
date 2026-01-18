package api

import (
	"net/http"
	"ponglockdown/server/api/handlers"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/v1/leaderboard", handlers.LeaderboardGet)
	mux.HandleFunc("/v1/leaderboard/submit", handlers.LeaderboardSubmit)
	mux.HandleFunc("/v1/replay/upload", handlers.ReplayUpload)
}
