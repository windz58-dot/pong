## What this patch does
- Adds TMP shadow banner Text object into Gameplay UI prefabs
- Wires `shadowBannerText` field automatically

## If the menu item doesn't find your prefab
Your UI might be generated at runtime or the prefab doesn't include GameplayUIController.
In that case:
- Put GameplayUIController on the prefab root or a child under Canvas
- Re-run installer
