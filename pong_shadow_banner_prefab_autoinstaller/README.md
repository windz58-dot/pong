# PONG — Shadow Banner Prefab Auto-Installer (TMP)

This follow-up patch adds a **Shadow Banner TMP_Text** into your existing **Gameplay UI prefab(s)** automatically.

It installs:
- Editor tool: `PONG/Install Shadow Banner (GameplayUIController prefabs)`
- Injects child: `ShadowBannerText` (TMP_Text) into any prefab containing `GameplayUIController`
- Auto-wires the serialized field `shadowBannerText` in `GameplayUIController` (from the fully-wired patch)

## Requirements
1) You already merged `pong_fully_wired_merge_patch.zip`
2) `GameplayUIController` class is marked as `partial` (per PATCH_NOTES.md)
3) TextMeshPro is installed (Package Manager)

## How to run
Unity Editor → Menu: **PONG → Install Shadow Banner (GameplayUIController prefabs)**

Result
- Prefabs are modified in-place (with Undo).
- The banner is disabled by default and only shows when Shadow Mode is enabled.

