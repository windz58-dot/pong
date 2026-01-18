using System;
using System.Collections;
using Pong.Net;
using UnityEngine;

namespace Pong.NetEx
{
    public sealed class LeaderboardClient : MonoBehaviour
    {
        [SerializeField] private BackendClient backend;

        private void Awake()
        {
            if (!backend) backend = FindObjectOfType<BackendClient>();
        }

        public IEnumerator Submit(string json, Action<string> ok, Action<string> err)
        {
            if (!backend) { err?.Invoke("BackendClient missing"); yield break; }
            yield return backend.PostJson("/v1/leaderboard/submit", json, ok, err);
        }
    }
}
