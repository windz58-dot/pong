package handlers

type RiskResult struct {
	RiskScore float64 `json:"risk_score"`
	HighRisk  bool    `json:"high_risk"`
}

type ReplayBlob struct {
	Format     string `json:"format"`
	PayloadB64 string `json:"payloadB64"`
	Sha256     string `json:"sha256"`
}

type LeaderboardSubmitReq struct {
	PlayerId      string   `json:"playerId"`
	SessionId     string   `json:"sessionId"`
	Season        string   `json:"season"`
	ScoreDelta    int      `json:"scoreDelta"`
	AttemptId     string   `json:"attemptId"`
	RewardEventId string   `json:"rewardEventId"`
	Replay        ReplayBlob `json:"replay"`
	Device        struct {
		DeviceId   string `json:"deviceId"`
		Platform   string `json:"platform"`
		AppVersion string `json:"appVersion"`
	} `json:"device"`
	Net struct {
		IpHint  string `json:"ipHint"`
		AsnHint string `json:"asnHint"`
	} `json:"net"`
}

type LeaderboardSubmitResp struct {
	Accepted bool      `json:"accepted"`
	NewRank  int       `json:"newRank"`
	NewScore int       `json:"newScore"`
	Risk     RiskResult `json:"risk"`
	Shadow   bool      `json:"shadow"`
	Reason   string    `json:"reason"`
}

type LeaderboardItem struct {
	Rank        int    `json:"rank"`
	PlayerId    string `json:"playerId"`
	DisplayName string `json:"displayName"`
	Score       int    `json:"score"`
	UpdatedAt   string `json:"updatedAt"`
}

type LeaderboardGetResp struct {
	Season   string          `json:"season"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
	Items    []LeaderboardItem `json:"items"`
}
