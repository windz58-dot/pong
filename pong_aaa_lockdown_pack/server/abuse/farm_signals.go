package abuse

type FarmSignals struct {
	AccountsPerDevice7d float64
	AccountsPerIP24h    float64
	AttemptsPerMin      float64
	NewAccount          bool
}

type FarmResult struct {
	Score01 float64
	Shadow  bool
	Reason  string
}

func FarmScore01(s FarmSignals) FarmResult {
	a := 0.0
	if s.NewAccount { a = 0.25 }
	b := clamp01((s.AccountsPerDevice7d - 2) / 8)
	c := clamp01((s.AccountsPerIP24h - 3) / 12)
	d := clamp01((s.AttemptsPerMin - 0.6) / 2.0)

	score := clamp01(0.25*a + 0.30*b + 0.25*c + 0.20*d)
	shadow := score >= 0.65 || (b >= 0.85) || (d >= 0.9)
	reason := ""
	if shadow { reason = "farm_throttle" }
	return FarmResult{Score01: score, Shadow: shadow, Reason: reason}
}

func clamp01(x float64) float64 {
	if x < 0 { return 0 }
	if x > 1 { return 1 }
	return x
}
