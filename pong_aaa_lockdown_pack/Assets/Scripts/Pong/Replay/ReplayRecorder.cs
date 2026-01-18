using System;
using System.Security.Cryptography;
using System.Text;
using Pong.Gameplay;
using UnityEngine;

namespace Pong.Replay
{
    public sealed class ReplayRecorder : MonoBehaviour
    {
        [SerializeField] private BallController ball;
        [SerializeField] private ShotTracker shotTracker;
        [SerializeField] private int maxFrames = 240;
        [SerializeField] private float frameInterval = 0.033f;

        private ReplayPayload _payload;
        private float _t0;
        private float _next;
        private bool _recording;

        public string AttemptId { get; private set; }

        public void Begin(string attemptId, Vector2 aimXZ, float power01)
        {
            AttemptId = attemptId;
            if (!ball) ball = FindObjectOfType<BallController>();
            if (!shotTracker) shotTracker = FindObjectOfType<ShotTracker>();

            _payload = new ReplayPayload { attemptId = attemptId };
            _payload.strike.aimX = Mathf.Clamp(aimXZ.x, -1f, 1f);
            _payload.strike.aimZ = Mathf.Clamp(aimXZ.y, -1f, 1f);
            _payload.strike.power01 = Mathf.Clamp01(power01);

            _payload.frames.Clear();
            _t0 = Time.time;
            _next = _t0;
            _recording = true;
        }

        public void End(bool scored, int potIndex)
        {
            if (_payload == null) return;
            _recording = false;

            _payload.summary.scored = scored;
            _payload.summary.potIndex = potIndex;

            if (shotTracker)
            {
                _payload.summary.maxSpeed = shotTracker.MaxSpeed;
                _payload.summary.flightSeconds = shotTracker.FlightSeconds;
                _payload.summary.bounces = shotTracker.Bounces;
            }
        }

        private void Update()
        {
            if (!_recording || _payload == null || !ball || ball.RB == null) return;
            if (_payload.frames.Count >= maxFrames) { _recording = false; return; }

            if (Time.time >= _next)
            {
                _next += frameInterval;
                var rb = ball.RB;
                var t = Time.time - _t0;

                _payload.frames.Add(new ReplayFrame{
                    t=t,
                    px=ball.transform.position.x, py=ball.transform.position.y, pz=ball.transform.position.z,
                    vx=rb.velocity.x, vy=rb.velocity.y, vz=rb.velocity.z
                });
            }
        }

        public string ToJson() => JsonUtility.ToJson(_payload);

        public static string Sha256Hex(string s)
        {
            using var sha = SHA256.Create();
            var b = sha.ComputeHash(Encoding.UTF8.GetBytes(s));
            var sb = new StringBuilder(b.Length*2);
            foreach (var x in b) sb.Append(x.ToString("x2"));
            return sb.ToString();
        }
    }
}
