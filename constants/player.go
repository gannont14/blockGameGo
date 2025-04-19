package constants

const (
  // world and player related
  RenderDistance = 3
  PlayerMoveSpeed = 0.2
  PlayerMouseSensitivity = 0.1
  PlayerHeight = 3
  PlayerHighlightRange = 10.0

  // inventory related
  PlayerInventoryRows = 4
  PlayerInventoryCols = 9
  PlayerInventoryWidth = 500
  PlayerInventorySlotMargin = 3

  // HUD Related
  // A good number would have (width - margin(numSlots + 1)) / numSlots
  //          be a relatively even number
  HUDHotbarWidth = 698 // ... thus 698
  HUDHotbarYOffset = 90
  HUDHotbarSlotMargin = 5
)
