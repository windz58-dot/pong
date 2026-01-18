package storage

import (
	"sort"
	"sync"
	"time"
)

type MemoryStore struct {
	mu sync.Mutex
	scores  map[string]map[string]int
	updated map[string]map[string]time.Time
	replays map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		scores: make(map[string]map[string]int),
		updated: make(map[string]map[string]time.Time),
		replays: make(map[string]string),
	}
}

func (m *MemoryStore) AddScore(season, playerId string, delta int) (int, error) {
	m.mu.Lock(); defer m.mu.Unlock()
	if _, ok := m.scores[season]; !ok {
		m.scores[season] = make(map[string]int)
		m.updated[season] = make(map[string]time.Time)
	}
	m.scores[season][playerId] += delta
	m.updated[season][playerId] = time.Now().UTC()
	return m.scores[season][playerId], nil
}

func (m *MemoryStore) GetRank(season, playerId string) (int, int, error) {
	m.mu.Lock(); defer m.mu.Unlock()
	entries := m.collect(season)
	sort.Slice(entries, func(i, j int) bool { return entries[i].Score > entries[j].Score })
	for i, e := range entries {
		if e.PlayerId == playerId {
			return i+1, e.Score, nil
		}
	}
	return 0, 0, nil
}

func (m *MemoryStore) Top(season string, page, pageSize int) ([]Entry, error) {
	m.mu.Lock(); defer m.mu.Unlock()
	entries := m.collect(season)
	sort.Slice(entries, func(i, j int) bool { return entries[i].Score > entries[j].Score })
	start := page * pageSize
	if start >= len(entries) { return []Entry{}, nil }
	end := start + pageSize
	if end > len(entries) { end = len(entries) }
	return entries[start:end], nil
}

func (m *MemoryStore) RecordReplay(attemptId, replayId, sha256 string, createdAt time.Time) error {
	m.mu.Lock(); defer m.mu.Unlock()
	m.replays[attemptId] = sha256
	return nil
}

func (m *MemoryStore) collect(season string) []Entry {
	var out []Entry
	for pid, sc := range m.scores[season] {
		out = append(out, Entry{PlayerId: pid, Score: sc, UpdatedAt: m.updated[season][pid]})
	}
	return out
}
