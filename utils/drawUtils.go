package utils

import (
	"blockProject/constants"
	types "blockProject/types"
	. "github.com/gen2brain/raylib-go/raylib"
)

func DrawCrosshair() {
  // this is straight from claude, not going to lie
    // Get screen center
    screenWidth := GetScreenWidth()
    screenHeight := GetScreenHeight()
    centerX := screenWidth / 2
    centerY := screenHeight / 2
    
    // Crosshair settings
    crosshairSize := 10
    crosshairThickness := 2
    crosshairColor := Black
    
    // Draw horizontal line
    DrawLineEx(
        Vector2{X: float32(centerX - crosshairSize), Y: float32(centerY)},
        Vector2{X: float32(centerX + crosshairSize), Y: float32(centerY)},
        float32(crosshairThickness),
        crosshairColor,
    )
    
    // Draw vertical line
    DrawLineEx(
        Vector2{X: float32(centerX), Y: float32(centerY - crosshairSize)},
        Vector2{X: float32(centerX), Y: float32(centerY + crosshairSize)},
        float32(crosshairThickness),
        crosshairColor,
    )
}

func drawBlock(b types.Block){
  // find center point of block
  centerPoint := b.CenterPoint()

  c := b.BlockColor()
  if c != White{
    DrawCube(centerPoint,
      constants.BlockWidth,
      constants.BlockHeight,
      constants.BlockDepth,
      b.BlockColor())
    // Draw the bounding box to debug
    // DrawBoundingBox(b.BoundBox, Blue)
  }

  if b.Focused{
    DrawCubeWires(centerPoint,
      constants.BlockWidth,
      constants.BlockHeight,
      constants.BlockDepth,
      Black)
  }
}

func drawChunk(c types.Chunk){
  if(len(c.Blocks)) == 0{
    return 
  }
  count := 0
  for i := range (len(c.Blocks)){
    for j := range(len(c.Blocks[0])){
      for k := range(len(c.Blocks[0][0])){
        b := c.Blocks[i][j][k]
        drawBlock(b)
        count += 1
      }
    }
  }

  // DEBUG, draw borde around chunk
  DrawChunkBorder(c)

  // fmt.Println("Drew Chunks: ", count)
}

func DrawChunks(w types.World, p types.Player, idcs []types.ChunkIndex) {

  // moved out to main function
  // idcs := types.GetRenderableChunkIndeces(p, w)

  if len(idcs) == 0{
    return
  }

  // new and improved
  for _, val := range idcs{
    j, i := val.UnboxChunkIndex()
    c := w.Chunks[i][j]
    drawChunk(c)
  }
}

func DrawInventory(inv types.Inventory) {
  return
}


/*
    this is WAY WAY WAY overengineered to be flexible with 
    different widths of the hotbar, as well as number of items

*/
func DrawHotbar(inv types.Inventory) {
  // calculate the required values, such as height, which 
  // would update the slots height and width

  l := len(inv.Slots)
  numHBSlots := constants.PlayerInventoryCols
  hotbarItemStacks := inv.Slots[(l - numHBSlots): l ]

  // figure out how many margins
  numMargins := numHBSlots + 1

  // How much space total the margins take up
  xMarginSpace := (numMargins * constants.HUDHotbarSlotMargin)
  // slot height/width
  slotSize := (constants.HUDHotbarWidth - xMarginSpace) / numHBSlots
  // find hotbar height now based on that
  HUDHotbarHeight := slotSize + (2 * constants.HUDHotbarSlotMargin)

  origX := int32(GetScreenWidth()/2) - (constants.HUDHotbarWidth/2)
  origY := int32(GetScreenHeight()) - constants.HUDHotbarYOffset

  // background
  DrawRectangle(origX,
    origY,
    constants.HUDHotbarWidth,
    int32(HUDHotbarHeight),
    White,
    )

  // now draw the slots
  DrawHotBarSlots(hotbarItemStacks, slotSize, int(origX), int(origY))
}

func DrawHotBarSlots(slots []types.ItemStack, slotSize int, origX int, origY int) {
  // should be able to take variable amount of items,
  // and display the user's hot bar items
  yOffset := constants.HUDHotbarSlotMargin
  for i:=range slots{
    xOffset := (i * (slotSize + constants.HUDHotbarSlotMargin)) + constants.HUDHotbarSlotMargin
    DrawHotBarSlot(slots[i], NewVector2(float32(origX + xOffset), float32(origY + yOffset)), slotSize)
  }
}

func DrawHotBarSlot(is types.ItemStack, pos Vector2, slotSize int) {
  DrawRectangle(
    int32(pos.X), 
    int32(pos.Y),
    int32(slotSize),
    int32(slotSize),
    Gray,
    )
}


