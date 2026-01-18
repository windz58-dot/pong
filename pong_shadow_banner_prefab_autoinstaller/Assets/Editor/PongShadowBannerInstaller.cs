using System.Linq;
using UnityEditor;
using UnityEngine;
using UnityEngine.UI;
using TMPro;

public static class PongShadowBannerInstaller
{
    private const string ChildName = "ShadowBannerText";

    [MenuItem("PONG/Install Shadow Banner (GameplayUIController prefabs)")]
    public static void Install()
    {
        // Find all prefabs
        string[] guids = AssetDatabase.FindAssets("t:Prefab");
        int touched = 0;
        int already = 0;

        foreach (var g in guids)
        {
            string path = AssetDatabase.GUIDToAssetPath(g);
            if (string.IsNullOrEmpty(path)) continue;

            // Quick filter: load prefab root and check for GameplayUIController
            var prefabRoot = PrefabUtility.LoadPrefabContents(path);
            try
            {
                var ui = prefabRoot.GetComponentInChildren<Pong.Gameplay.GameplayUIController>(true);
                if (ui == null) continue;

                // Ensure a Canvas exists in this prefab subtree (UI prefabs typically do)
                var canvas = prefabRoot.GetComponentInChildren<Canvas>(true);
                if (canvas == null)
                {
                    Debug.LogWarning($"[PONG] Skipped (no Canvas): {path}");
                    continue;
                }

                // If exists already
                var existing = prefabRoot.GetComponentsInChildren<TMP_Text>(true)
                    .FirstOrDefault(t => t != null && t.gameObject.name == ChildName);

                if (existing != null)
                {
                    // Try to auto-wire if missing
                    TryWire(ui, existing);
                    already++;
                    continue;
                }

                // Choose parent: prefer a top-level UI container under canvas
                Transform parent = ui.transform;
                var layout = ui.GetComponentInChildren<VerticalLayoutGroup>(true);
                if (layout != null) parent = layout.transform;
                else
                {
                    var panel = ui.GetComponentsInChildren<RectTransform>(true)
                        .FirstOrDefault(rt => rt.name.ToLower().Contains("panel"));
                    if (panel != null) parent = panel;
                }

                var go = new GameObject(ChildName, typeof(RectTransform), typeof(TMP_Text));
                go.transform.SetParent(parent, false);

                var rt = go.GetComponent<RectTransform>();
                // Anchor top-center
                rt.anchorMin = new Vector2(0.5f, 1f);
                rt.anchorMax = new Vector2(0.5f, 1f);
                rt.pivot = new Vector2(0.5f, 1f);
                rt.anchoredPosition = new Vector2(0f, -24f);
                rt.sizeDelta = new Vector2(900f, 90f);

                var txt = go.GetComponent<TMP_Text>();
                txt.text = "Reward delivered (verified).";
                txt.enableWordWrapping = false;
                txt.alignment = TextAlignmentOptions.Center;
                txt.fontSize = 34;
                txt.fontStyle = FontStyles.Bold;
                txt.raycastTarget = false;

                // Subtle glow (works with TMP default material; safe even if no URP bloom)
                txt.outlineWidth = 0.25f;
                txt.outlineColor = new Color(0f, 0f, 0f, 0.65f);

                // Disabled by default
                go.SetActive(false);

                TryWire(ui, txt);

                PrefabUtility.SaveAsPrefabAsset(prefabRoot, path);
                touched++;
                Debug.Log($"[PONG] Installed Shadow Banner into: {path}");
            }
            finally
            {
                PrefabUtility.UnloadPrefabContents(prefabRoot);
            }
        }

        AssetDatabase.SaveAssets();
        AssetDatabase.Refresh();

        EditorUtility.DisplayDialog("PONG Shadow Banner Installer",
            $"Done. Modified: {touched} prefab(s). Already had banner: {already}.", "OK");
    }

    private static void TryWire(Pong.Gameplay.GameplayUIController ui, TMP_Text txt)
    {
        if (ui == null || txt == null) return;

        // Serialized field from GameplayUIController_ShadowExtensions.cs:
        // [SerializeField] private TMPro.TMP_Text shadowBannerText;
        var so = new SerializedObject(ui);
        var prop = so.FindProperty("shadowBannerText");
        if (prop != null && prop.objectReferenceValue == null)
        {
            prop.objectReferenceValue = txt;
            so.ApplyModifiedPropertiesWithoutUndo();
        }
    }
}
