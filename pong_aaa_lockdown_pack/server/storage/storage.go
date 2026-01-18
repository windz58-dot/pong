package storage

import "time"

type LeaderboardStore interface {
	AddScore(season, playerId string, delta int) (newScore int, err error)
	GetRank(season, playerId string) (rank int, score int, err error)
	Top(season string, page, pageSize int) ([]Entry, error)
	RecordReplay(attemptId, replayId, sha256 string, createdAt time.Time) error
}

type Entry struct {
	PlayerId  string
	Score     int
	UpdatedAt time.Time
}
