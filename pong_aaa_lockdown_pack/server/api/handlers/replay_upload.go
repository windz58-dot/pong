package handlers

import (
	"net/http"
	"time"
	"ponglockdown/server/replay"
)

type ReplayUploadReq struct {
	PlayerId  string   `json:"playerId"`
	AttemptId string   `json:"attemptId"`
	Replay    ReplayBlob `json:"replay"`
}
type ReplayUploadResp struct { Ok bool `json:"ok"`; ReplayId string `json:"replayId"` }

func ReplayUpload(w http.ResponseWriter, r *http.Request) {
	var req ReplayUploadReq
	if err := readJSON(r, &req); err != nil { writeJSON(w, 400, map[string]any{"error":"bad_json"}); return }

	rep, err := replay.ParseAndValidate(req.Replay.PayloadB64, req.Replay.Sha256)
	if err != nil || rep.AttemptId != req.AttemptId {
		writeJSON(w, 400, map[string]any{"error":"invalid_replay"}); return
	}
	replayId := "r_" + req.AttemptId
	_ = Store.RecordReplay(req.AttemptId, replayId, req.Replay.Sha256, time.Now().UTC())
	writeJSON(w, 200, ReplayUploadResp{Ok:true, ReplayId:replayId})
}
