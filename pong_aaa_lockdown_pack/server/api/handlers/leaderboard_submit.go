package handlers

import (
	"net/http"
	"time"

	"ponglockdown/server/abuse"
	"ponglockdown/server/replay"
)

func LeaderboardSubmit(w http.ResponseWriter, r *http.Request) {
	var req LeaderboardSubmitReq
	if err := readJSON(r, &req); err != nil {
		writeJSON(w, 400, map[string]any{"error":"bad_json"}); return
	}

	if !deviceBucket.Allow(1) || !ipBucket.Allow(1) {
		writeJSON(w, 429, LeaderboardSubmitResp{Accepted:false, Shadow:true, Reason:"throttled", Risk: RiskResult{RiskScore:85, HighRisk:true}})
		return
	}

	rep, err := replay.ParseAndValidate(req.Replay.PayloadB64, req.Replay.Sha256)
	if err != nil || rep.AttemptId != req.AttemptId {
		writeJSON(w, 400, LeaderboardSubmitResp{Accepted:false, Shadow:true, Reason:"invalid_replay", Risk: RiskResult{RiskScore:90, HighRisk:true}})
		return
	}

	// TODO: compute real farm signals from Redis/Postgres
	farm := abuse.FarmScore01(abuse.FarmSignals{
		AccountsPerDevice7d: 3,
		AccountsPerIP24h:    4,
		AttemptsPerMin:      0.8,
		NewAccount:          false,
	})
	shadow := farm.Shadow
	risk := RiskResult{RiskScore: 100 * farm.Score01, HighRisk: shadow}

	delta := req.ScoreDelta
	if !rep.Summary.Scored { delta = 0 } // server-authoritative
	newScore, _ := Store.AddScore(req.Season, req.PlayerId, delta)
	rank, _, _ := Store.GetRank(req.Season, req.PlayerId)

	writeJSON(w, 200, LeaderboardSubmitResp{
		Accepted:true, NewRank:rank, NewScore:newScore, Risk:risk, Shadow:shadow, Reason:farm.Reason,
	})

	_ = Store.RecordReplay(req.AttemptId, "r_"+req.AttemptId, req.Replay.Sha256, time.Now().UTC())
}
