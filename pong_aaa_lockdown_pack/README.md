# PONG — AAA Lock-down Pack (Fair & Square)

Adds:
✅ Server-side authoritative leaderboard + replay validation
✅ Device + account farm throttling
✅ Anomaly-driven shadow rewards for flagged users (still playable, no economy drain)

Layer on top of:
- pong_production_reward_engine_pack.zip
- pong_reward_engine_unity_integration_patch.zip OR pong_fair_square_full.zip

## Contents
Server (Go):
- /server/api : leaderboard + replay endpoints
- /server/replay : replay schema + validation
- /server/abuse : throttling + farm scoring
- /server/rewards_shadow : shadow reward downgrade mapper
- /server/storage : storage interfaces + in-memory store for testing

Unity:
- /Assets/Scripts/Pong/Replay : replay recorder + submit builder
- /Assets/Scripts/Pong/Net    : Leaderboard client
- /Assets/Scripts/Pong/Telemetry : minimal device signals helper

## Guarantees
- Leaderboard entries are accepted only with a valid replay payload.
- Server validates plausibility (format + bounds + monotonic frames).
- Farm/throttle signals can flip `shadow=true`; economy rewards downgrade server-side.
