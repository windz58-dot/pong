using System.Collections.Generic;
using UnityEngine;

namespace Pong.Telemetry
{
    public static class DeviceSignals
    {
        public static Dictionary<string, object> MinimalSignals()
        {
            return new Dictionary<string, object>
            {
                {"platform", Application.platform.ToString()},
                {"appVersion", Application.version},
                {"deviceModel", SystemInfo.deviceModel},
                {"memGB", Mathf.RoundToInt(SystemInfo.systemMemorySize/1024f)},
            };
        }
    }
}
