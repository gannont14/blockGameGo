package utils

import (
	"blockProject/constants"
	types "blockProject/types"
	"fmt"

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
    DrawItemSlot(slots[i], NewVector2(float32(origX + xOffset), float32(origY + yOffset)), slotSize, Gray)
  }
}

func DrawItemSlot(is types.ItemStack, pos Vector2, slotSize int, col Color) {
  DrawRectangle(
    int32(pos.X), 
    int32(pos.Y),
    int32(slotSize),
    int32(slotSize),
    col,
    )
}

func DrawInventory(inv types.Inventory){
  // divider between inventory and hotbar items
  dividerHeight := constants.PlayerInventorySlotMargin * 3.0
  // total margin space on y axis (numRows + 1)
  marginSpaceY := (constants.PlayerInventoryRows * constants.PlayerInventorySlotMargin) + dividerHeight
  marginSpaceX := constants.PlayerInventoryCols * constants.PlayerInventorySlotMargin
  
  // slot size 
  slotSize := (constants.PlayerInventoryWidth - (marginSpaceX)) / constants.PlayerInventoryCols

  // calculate what the height should be
  PlayerInventoryHeight := int(marginSpaceY) + (slotSize * constants.PlayerInventoryRows)

  // calculate the origin
  origX := (int32(GetScreenWidth()) / 2) - (constants.PlayerInventoryWidth / 2)
  origY := (int32(GetScreenHeight()) / 2) - (int32(PlayerInventoryHeight) / 2)

  bgOrigX := origX - constants.PlayerInventorySlotMargin
  bgOrigY := origY - constants.PlayerInventorySlotMargin

  // Draw the background
  DrawRectangle(bgOrigX , bgOrigY,
    constants.PlayerInventoryWidth + (2 * constants.PlayerInventorySlotMargin),
    int32(PlayerInventoryHeight) + (2 * constants.PlayerInventorySlotMargin),
    White)

  // Draw the individual slots
  DrawInventorySlots(inv.Slots, int(origX), int(origY), slotSize, int(dividerHeight))
}

func DrawInventorySlots(items []types.ItemStack, origX int, origY int, slotSize int, dividerHeight int){

  // figure out which slot the user is hovering over
  activeSlot := -1

  for i := range items {
    r, c := slotIndexToCoord(i)
    xOffset := (c * (slotSize + constants.PlayerInventorySlotMargin)) + constants.PlayerInventorySlotMargin
    yOffset := (r * (slotSize + constants.PlayerInventorySlotMargin)) + constants.PlayerInventorySlotMargin

    // for vals under the divider (hotbar item)
    if r == (constants.PlayerInventoryRows - 1){
      yOffset += dividerHeight
    }
    col := Gray
    origVec := NewVector2(float32(origX + xOffset), float32(origY + yOffset))
    // check if mouse in that square
    if pointInRectangle(GetMousePosition(), origVec, slotSize, slotSize){
      activeSlot = i 
      col = LightGray
    }

    DrawItemSlot(items[i], origVec, slotSize, col)
    drawItemStack(origVec, items[i], slotSize)
  }

  types.UpdateActiveSelectedSlot(activeSlot)
}

// will always be square so the size can just be one value
func drawItemStack(pos Vector2, is types.ItemStack, size int){
  if is.Count == 0{
    return
  }
  // for now will draw the item ID and then the count
  itemName := fmt.Sprint(is.Item.GetName())
  itemCount := fmt.Sprint(is.Count)

  // Draw the item's ID in the center
  DrawText(itemName, int32(pos.X), int32(pos.Y), 12, Blue)

  // Daw the item's count
  DrawText(itemCount, int32(pos.X), int32(pos.Y), 25, Red)


}


func slotIndexToCoord(ind int) (int, int) {
  // should never go over the amount of slots
  row := int(ind / constants.PlayerInventoryCols)
  col := ind % constants.PlayerInventoryCols

  fmt.Printf("Ind: %d, Row: %d, Col: %d\n",ind, row, col)

  return row, col
}

func pointInRectangle(point Vector2, orig Vector2, height int, width int) bool {
  // check x 
  px := int(point.X) // annoying to type cast every time
  ox := int(orig.X) // annoying to type cast every time
  checkX := (px > ox) && (px < (ox + width))
  
  // check y
  py := int(point.Y) // annoying to type cast every time
  oy := int(orig.Y) // annoying to type cast every time
  checkY := (py > oy) && (py < (oy + height))

  return checkX && checkY
}







