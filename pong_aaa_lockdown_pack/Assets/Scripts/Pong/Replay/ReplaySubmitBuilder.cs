using System;
using System.Text;
using Pong.Profile;
using Pong.Telemetry;
using UnityEngine;

namespace Pong.Replay
{
    [Serializable]
    public sealed class LeaderboardSubmitRequest
    {
        public string playerId;
        public string sessionId;
        public string season;
        public int scoreDelta;
        public string attemptId;
        public string rewardEventId;
        public ReplayBlob replay = new ReplayBlob();
        public Device device = new Device();
        public Net net = new Net();

        [Serializable] public sealed class ReplayBlob { public string format; public string payloadB64; public string sha256; }
        [Serializable] public sealed class Device { public string deviceId; public string platform; public string appVersion; }
        [Serializable] public sealed class Net { public string ipHint; public string asnHint; }
    }

    public static class ReplaySubmitBuilder
    {
        public static string Build(string sessionId, string season, int scoreDelta, string attemptId, string rewardEventId, string replayJson)
        {
            var p = PlayerProfile.LoadOrCreate();
            string b64 = Convert.ToBase64String(Encoding.UTF8.GetBytes(replayJson));
            string sha = ReplayRecorder.Sha256Hex(replayJson);

            var req = new LeaderboardSubmitRequest
            {
                playerId = p.playerId,
                sessionId = sessionId,
                season = season,
                scoreDelta = scoreDelta,
                attemptId = attemptId,
                rewardEventId = rewardEventId,
            };
            req.replay.format = "pong_replay_v1";
            req.replay.payloadB64 = b64;
            req.replay.sha256 = sha;

            req.device.deviceId = SystemInfo.deviceUniqueIdentifier;
            req.device.platform = Application.platform.ToString();
            req.device.appVersion = Application.version;

            req.net.ipHint = "";
            req.net.asnHint = "";

            Pong.Telemetry.TelemetryHub.I?.Track("leaderboard_submit_build", DeviceSignals.MinimalSignals());
            return JsonUtility.ToJson(req);
        }
    }
}
