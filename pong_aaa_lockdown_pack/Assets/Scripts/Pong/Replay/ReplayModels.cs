using System;
using System.Collections.Generic;

namespace Pong.Replay
{
    [Serializable] public sealed class ReplayFrame { public float t; public float px,py,pz; public float vx,vy,vz; }

    [Serializable]
    public sealed class ReplayPayload
    {
        public string format = "pong_replay_v1";
        public string attemptId;
        public Strike strike = new Strike();
        public List<ReplayFrame> frames = new List<ReplayFrame>();
        public Summary summary = new Summary();

        [Serializable] public sealed class Strike { public float aimX; public float aimZ; public float power01; }
        [Serializable] public sealed class Summary { public float maxSpeed; public float flightSeconds; public int bounces; public bool scored; public int potIndex; }
    }
}
