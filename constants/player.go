package constants

const (
  // world and player related
  RenderDistance = 3
  PlayerMoveSpeed = 0.2
  PlayerMouseSensitivity = 0.06
  PlayerHeight = 3
  PlayerHighlightRange = 10.0

  // inventory related
  PlayerInventoryRows = 4
  PlayerInventoryCols = 9
  PlayerInventoryWidth = 900
  PlayerInventorySlotMargin = 3

  HotbarOffset = PlayerInventoryCols * (PlayerInventoryRows - 1)

  // HUD Related
  // A good number would have (width - margin(numSlots + 1)) / numSlots
  //          be a relatively even number
  HUDHotbarWidth = 698 // ... thus 698
  HUDHotbarYOffset = 90
  HUDHotbarSlotMargin = 5
  HUDHotbarBorderThickness = 5

	// Screen size
	ScreenWidth = 1600
	ScreenHeight = 900
)
