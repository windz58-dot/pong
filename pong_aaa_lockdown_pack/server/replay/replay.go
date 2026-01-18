package replay

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math"
)

type Frame struct {
	T  float32 `json:"t"`
	Px float32 `json:"px"`
	Py float32 `json:"py"`
	Pz float32 `json:"pz"`
	Vx float32 `json:"vx"`
	Vy float32 `json:"vy"`
	Vz float32 `json:"vz"`
}

type Replay struct {
	Format   string `json:"format"`
	AttemptId string `json:"attemptId"`
	Strike struct {
		AimX float32 `json:"aimX"`
		AimZ float32 `json:"aimZ"`
		Power01 float32 `json:"power01"`
	} `json:"strike"`
	Frames []Frame `json:"frames"`
	Summary struct {
		MaxSpeed float32 `json:"maxSpeed"`
		FlightSeconds float32 `json:"flightSeconds"`
		Bounces int `json:"bounces"`
		Scored bool `json:"scored"`
		PotIndex int `json:"potIndex"`
	} `json:"summary"`
}

func DecodeB64(b64 string) ([]byte, error) { return base64.StdEncoding.DecodeString(b64) }
func Sha256Hex(b []byte) string { s:=sha256.Sum256(b); return hex.EncodeToString(s[:]) }

func ParseAndValidate(payloadB64, expectedSha string) (*Replay, error) {
	raw, err := DecodeB64(payloadB64)
	if err != nil { return nil, err }
	if expectedSha != "" && Sha256Hex(raw) != expectedSha {
		return nil, errors.New("sha256 mismatch")
	}
	var r Replay
	if err := json.Unmarshal(raw, &r); err != nil { return nil, err }
	if r.Format != "pong_replay_v1" { return nil, errors.New("bad format") }
	if len(r.Frames) < 8 || len(r.Frames) > 400 { return nil, errors.New("bad frame count") }
	if r.Strike.Power01 < 0 || r.Strike.Power01 > 1 { return nil, errors.New("power01 out of range") }
	if math.Abs(float64(r.Strike.AimX)) > 1.01 || math.Abs(float64(r.Strike.AimZ)) > 1.01 { return nil, errors.New("aim out of range") }

	prevT := float32(-1)
	for _, f := range r.Frames {
		if f.T < prevT { return nil, errors.New("non-monotonic time") }
		prevT = f.T
	}
	if r.Summary.FlightSeconds < 0 || r.Summary.FlightSeconds > 30 { return nil, errors.New("flightSeconds out of range") }
	if r.Summary.MaxSpeed > 25 { return nil, errors.New("maxSpeed too high") }
	if r.Summary.Bounces < 0 || r.Summary.Bounces > 35 { return nil, errors.New("bounces too high") }
	return &r, nil
}
