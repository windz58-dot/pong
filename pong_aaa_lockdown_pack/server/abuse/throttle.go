package abuse

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mu sync.Mutex
	tokens float64
	rate float64
	cap float64
	last time.Time
}

func NewTokenBucket(ratePerMin, cap float64) *TokenBucket {
	return &TokenBucket{tokens: cap, rate: ratePerMin/60.0, cap: cap, last: time.Now()}
}

func (b *TokenBucket) Allow(cost float64) bool {
	b.mu.Lock(); defer b.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(b.last).Seconds()
	b.last = now
	b.tokens = min(b.cap, b.tokens + elapsed*b.rate)
	if b.tokens >= cost {
		b.tokens -= cost
		return true
	}
	return false
}
func min(a,b float64) float64 { if a<b { return a }; return b }
