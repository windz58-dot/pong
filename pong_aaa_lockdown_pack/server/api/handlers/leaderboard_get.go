package handlers

import (
	"net/http"
	"strconv"
	"time"
)

func LeaderboardGet(w http.ResponseWriter, r *http.Request) {
	season := r.URL.Query().Get("season")
	if season == "" { season = "S1" }
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if pageSize <= 0 || pageSize > 100 { pageSize = 50 }
	if page < 0 { page = 0 }

	entries, _ := Store.Top(season, page, pageSize)
	items := make([]LeaderboardItem, 0, len(entries))
	for i, e := range entries {
		items = append(items, LeaderboardItem{
			Rank: page*pageSize + i + 1,
			PlayerId: e.PlayerId,
			DisplayName: "Player",
			Score: e.Score,
			UpdatedAt: e.UpdatedAt.Format(time.RFC3339),
		})
	}

	writeJSON(w, 200, LeaderboardGetResp{Season:season, Page:page, PageSize:pageSize, Items:items})
}
