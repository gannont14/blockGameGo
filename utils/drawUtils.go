package utils

import (
	"blockProject/constants"
	gamestate "blockProject/gamestate"
	"blockProject/textures"
	"blockProject/types"
	"fmt"

	// "fmt"

	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	gs *gamestate.GameState
)

func drawItemDurability(item types.Item, pos Vector2, width, height int) {
	// check if item is tool or somethign with durability
	if item == nil {
		return 
	}

	if gs == nil {
		gs = gamestate.Get()
		return 
	}

	_, isTool := item.(*types.ToolItem)
	if !isTool {
		return 
	}

	maxDur := item.(*types.ToolItem).MaxDurability
	dur := item.(*types.ToolItem).Durability
	var durPercent float32 = float32(dur) / float32(maxDur)

	DrawProgressBar(pos, height, width, float32(durPercent), DarkGray, Green)
}

func DrawItemName(pos Vector2, fontSize int) {
	p := gamestate.Get().Player
	activeItem, _ := p.GetActiveItem()

	// nothign in hand
	if activeItem == nil {
		return
	}

	// grab the name of the active item
	name := activeItem.GetName()
	textWidth := MeasureText(name, int32(fontSize))
	centeredX := int32(pos.X - (float32(textWidth) / 2))

	drawLayeredText(Vector2{X: float32(centeredX), Y: pos.Y}, name, fontSize, Black, White)
}

func DrawProgressBar(pos Vector2, height int, width int, progress float32, bgColor Color, barColor Color) {
	startPos := NewVector2(pos.X-(float32(width)/2.0), pos.Y)
	endPos := NewVector2(pos.X+(float32(width)/2.0), pos.Y)

	// background
	DrawLineEx(startPos,
		endPos,
		float32(height),
		bgColor)

	// find what width should be
	progWidth := progress * float32(width)
	pStartPos := NewVector2(pos.X-(float32(width)/2.0), pos.Y)
	pEndPos := NewVector2(pStartPos.X+progWidth, pos.Y)

	// bar
	DrawLineEx(pStartPos,
		pEndPos,
		float32(height),
		barColor,
	)
}

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

func drawBlock(b types.Block) {
	// find center point of block
	centerPoint := b.CenterPoint()

	if b.Type != types.Air {

		// nilPos := textures.TextureAtlasPosition{}
		// if b.BlockItem.AtlasPosition ==  nilPos {
		// fallback
		// 	DrawCube(centerPoint,
		// 		constants.BlockWidth,
		// 		constants.BlockHeight,
		// 		constants.BlockDepth,
		// 		b.Color,
		// 		)
		// }

		// draw the textured cube
		drawTexturedCube(centerPoint,
			constants.BlockWidth,
			constants.BlockHeight,
			constants.BlockDepth,
			&b,
		)

		// Draw the bounding box to debug
		// DrawBoundingBox(b.BoundBox, Blue)
	}

	if b.Focused {
		DrawCubeWires(centerPoint,
			constants.BlockWidth,
			constants.BlockHeight,
			constants.BlockDepth,
			Black)
	}
}

func drawTexturedCube(pos Vector3, width, height, depth float32, b *types.Block) {
	ta := gamestate.Get().World.BlockAtlas
	row, col := b.BlockItem.AtlasPosition.Val()
	blockTexture := textures.NewBlockTexture(ta, row, col)

	// half of each
	hs_w := width / 2
	hs_h := height / 2
	hs_d := depth / 2

	// Define vertices of cube
	v000 := Vector3{X: pos.X - hs_w, Y: pos.Y - hs_h, Z: pos.Z - hs_d}
	v001 := Vector3{X: pos.X - hs_w, Y: pos.Y - hs_h, Z: pos.Z + hs_d}
	v010 := Vector3{X: pos.X - hs_w, Y: pos.Y + hs_h, Z: pos.Z - hs_d}
	v011 := Vector3{X: pos.X - hs_w, Y: pos.Y + hs_h, Z: pos.Z + hs_d}
	v100 := Vector3{X: pos.X + hs_w, Y: pos.Y - hs_h, Z: pos.Z - hs_d}
	v101 := Vector3{X: pos.X + hs_w, Y: pos.Y - hs_h, Z: pos.Z + hs_d}
	v110 := Vector3{X: pos.X + hs_w, Y: pos.Y + hs_h, Z: pos.Z - hs_d}
	v111 := Vector3{X: pos.X + hs_w, Y: pos.Y + hs_h, Z: pos.Z + hs_d}

	// Helper to draw a face with texture coordinates
	drawFace := func(a, b, c, d Vector3, coords [4]Vector2) {
		Begin(Quads)
		SetTexture(blockTexture.Atlas.Texture.ID)

		// Set vertex color to white to ensure full texture brightness
		Color4ub(255, 255, 255, 255)
		TexCoord2f(coords[0].X, coords[0].Y)
		Vertex3f(a.X, a.Y, a.Z)

		Color4ub(255, 255, 255, 255)
		TexCoord2f(coords[1].X, coords[1].Y)
		Vertex3f(b.X, b.Y, b.Z)

		Color4ub(255, 255, 255, 255)
		TexCoord2f(coords[2].X, coords[2].Y)
		Vertex3f(c.X, c.Y, c.Z)

		Color4ub(255, 255, 255, 255)
		TexCoord2f(coords[3].X, coords[3].Y)
		Vertex3f(d.X, d.Y, d.Z)

		End()
	}

	// Get the BlockTextureMap from the BlockTexture
	textureMap := blockTexture.BlockTexturemap

	// Draw each face with corresponding texture coordinates
	drawFace(v011, v111, v110, v010, textureMap["top"])    // Top face (Y+)
	drawFace(v000, v100, v101, v001, textureMap["bottom"]) // Bottom face (Y-)
	drawFace(v001, v101, v111, v011, textureMap["front"])  // Front face (Z+)
	drawFace(v100, v000, v010, v110, textureMap["back"])   // Back face (Z-)
	drawFace(v000, v001, v011, v010, textureMap["left"])   // Left face (X-)
	drawFace(v101, v100, v110, v111, textureMap["right"])  // Right face (X+)

}

func drawChunk(c types.Chunk) {
	if (len(c.Blocks)) == 0 {
		return
	}
	count := 0
	for i := range len(c.Blocks) {
		for j := range len(c.Blocks[0]) {
			for k := range len(c.Blocks[0][0]) {
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

func DrawChunks(idcs []types.ChunkIndex) {
	w := gamestate.Get().World

	// moved out to main function
	// idcs := types.GetRenderableChunkIndeces(p, w)

	if len(idcs) == 0 {
		return
	}

	// new and improved
	for _, val := range idcs {
		j, i := val.UnboxChunkIndex()
		c := w.Chunks[i][j]
		drawChunk(c)
	}
}

/*
this is WAY WAY WAY overengineered to be flexible with
different widths of the hotbar, as well as number of items
*/
func DrawHotbar(inv types.Inventory, activeItemSlot int) {
	// calculate the required values, such as height, which
	// would update the slots height and width

	l := len(inv.Slots)
	numHBSlots := constants.PlayerInventoryCols
	hotbarItemStacks := inv.Slots[(l - numHBSlots):l]

	// figure out how many margins
	numMargins := numHBSlots + 1

	// How much space total the margins take up
	xMarginSpace := (numMargins * constants.HUDHotbarSlotMargin)
	// slot height/width
	slotSize := (constants.HUDHotbarWidth - xMarginSpace) / numHBSlots
	// find hotbar height now based on that
	HUDHotbarHeight := slotSize + (2 * constants.HUDHotbarSlotMargin)

	origX := int32(GetScreenWidth()/2) - (constants.HUDHotbarWidth / 2)
	origY := int32(GetScreenHeight()) - constants.HUDHotbarYOffset

	// background
	DrawRectangle(origX,
		origY,
		constants.HUDHotbarWidth,
		int32(HUDHotbarHeight),
		White,
	)

	// now draw the slots
	DrawHotBarSlots(hotbarItemStacks, slotSize, int(origX), int(origY), activeItemSlot)
}

func DrawHotBarSlots(slots []types.ItemStack, slotSize int, origX int, origY int, activeSlot int) {
	// should be able to take variable amount of items,
	// and display the user's hot bar items
	yOffset := constants.HUDHotbarSlotMargin
	for i := range slots {
		hasBorder := false
		xOffset := (i * (slotSize + constants.HUDHotbarSlotMargin)) + constants.HUDHotbarSlotMargin

		// check if it is the player's active slot
		// fmt.Printf("Current: %d Active: %d\n", i, activeSlot)
		if i == activeSlot-constants.HotbarOffset {
			hasBorder = true
		}

		DrawItemSlot(slots[i], NewVector2(float32(origX+xOffset), float32(origY+yOffset)), slotSize, Gray, hasBorder)
		drawItemStack(NewVector2(float32(origX+xOffset), float32(origY+yOffset)), slots[i], slotSize)
	}
}

func DrawItemSlot(is types.ItemStack, pos Vector2, slotSize int, col Color, hasBorder bool) {
	DrawRectangle(
		int32(pos.X),
		int32(pos.Y),
		int32(slotSize),
		int32(slotSize),
		col,
	)
	if hasBorder {
		// draw the active border around the item
		rec := NewRectangle(
			pos.X-constants.HUDHotbarBorderThickness,
			pos.Y-constants.HUDHotbarBorderThickness,
			float32(slotSize+(2*constants.HUDHotbarBorderThickness)),
			float32(slotSize+(2*constants.HUDHotbarBorderThickness)),
		)
		DrawRectangleLinesEx(rec, constants.HUDHotbarBorderThickness, DarkGray)
	}
}

func DrawInventory(inv *types.Inventory) {
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
	DrawRectangle(bgOrigX, bgOrigY,
		constants.PlayerInventoryWidth+(2*constants.PlayerInventorySlotMargin),
		int32(PlayerInventoryHeight)+(2*constants.PlayerInventorySlotMargin),
		White)

	// Draw the individual slots
	DrawInventorySlots(inv, int(origX), int(origY), slotSize, int(dividerHeight))
}

func DrawInventorySlots(inv *types.Inventory, origX int, origY int, slotSize int, dividerHeight int) {
	items := inv.Slots

	// figure out which slot the user is hovering over
	activeSlot := -1

	for i := range items {
		r, c := slotIndexToCoord(i)
		xOffset := (c * (slotSize + constants.PlayerInventorySlotMargin)) + constants.PlayerInventorySlotMargin
		yOffset := (r * (slotSize + constants.PlayerInventorySlotMargin)) + constants.PlayerInventorySlotMargin

		// for vals under the divider (hotbar item)
		if r == (constants.PlayerInventoryRows - 1) {
			yOffset += dividerHeight
		}
		col := Gray
		origVec := NewVector2(float32(origX+xOffset), float32(origY+yOffset))
		// check if mouse in that square
		if pointInRectangle(GetMousePosition(), origVec, slotSize, slotSize) {
			activeSlot = i
			col = LightGray
		}

		DrawItemSlot(items[i], origVec, slotSize, col, false)
		drawItemStack(origVec, items[i], slotSize)
		// on click handler
	}

	types.UpdateActiveSelectedSlot(activeSlot)
	// after the draw call, make sure that there is an active slot
	if IsMouseButtonPressed(MouseLeftButton) {
		// update the inventory accordignly
		inv.HandleItemClick(activeSlot)
	}
	// PrintPlayerHand(inv)
	drawItemStack(GetMousePosition(), inv.Hand, 50)

	if activeSlot >= 0 {
		// draw the item name
		drawHoveringItemName(*inv,
			GetMousePosition(),
			activeSlot, 30)
	}
}

func drawHoveringItemName(inv types.Inventory, mousePos Vector2, activeSlot, fontSize int) {
	// figure out what item is in the players hand
	itemStack := inv.Slots[activeSlot]
	// empty
	if itemStack.Item == nil {
		return
	}

	// pull out item name
	name := itemStack.Item.GetName()
	nameWidth := MeasureText(name, int32(fontSize))

	textPosX, textPosY := mousePos.X + 3, mousePos.Y + float32(fontSize + 2)

	//  determine origin based on the whether or not it would stay on screen
	if (mousePos.X + float32(nameWidth)) > constants.ScreenWidth {
		// invert (subtract the text width and a few)
		textPosX -= float32(nameWidth + 3) 
	}

	// draw background now
	DrawRectangle(int32(textPosX),
		int32(textPosY),
		nameWidth + 10,
		int32(fontSize) + 2,
		DarkGray)

	// now draw the text over top
	drawLayeredText(Vector2{X: textPosX + 2, Y: textPosY + 2},
		name,
		fontSize, 
		Black,
		White)

}

// will always be square so the size can just be one value
func drawItemStack(pos Vector2, is types.ItemStack, size int) {
	ia := gamestate.Get().ItemAtlas
	if is.Item == nil {
		return
	}
	// for now will draw the item ID and then the count
	// itemName := fmt.Sprint(is.Item.GetName())
	itemCount := fmt.Sprint(is.Count)

	rec := ia.GetRectanglePos(is.Item.GetAtlasPosition())
	drawScaledRectangle(rec, pos, float32(size), ia)

	// Draw the item's ID in the center
	// DrawText(itemName, int32(pos.X), int32(pos.Y), 12, Blue)

	// Daw the item's count
	DrawText(itemCount, int32(pos.X), int32(pos.Y), 25, Red)

	// draw the items durability, in the bottom third
	durPos := Vector2{X: pos.X + float32(size/2), Y: pos.Y + float32(size* 9/10)}
	drawItemDurability(is.Item, durPos, int(float32(size) * float32(0.7)), 5)
}

func drawLayeredText(pos Vector2, text string, fontSize int, bgColor, fgColor Color) {
	// draw 'background' text, where it's black
	DrawText(fmt.Sprintf("%s", text), int32(pos.X-1), int32(pos.Y-1), int32(fontSize), bgColor)
	// draw 'foreground' text, where it's white
	DrawText(fmt.Sprintf("%s", text), int32(pos.X), int32(pos.Y), int32(fontSize), fgColor)

}

func drawScaledRectangle(tileRec Rectangle, pos Vector2, size float32, ia *textures.ItemAtlas) {
	scaledRec := NewRectangle(pos.X, pos.Y, float32(size), float32(size))

	DrawTexturePro(ia.Texture, tileRec, scaledRec, NewVector2(0, 0), 0.0, White)
}

func slotIndexToCoord(ind int) (int, int) {
	// should never go over the amount of slots
	row := int(ind / constants.PlayerInventoryCols)
	col := ind % constants.PlayerInventoryCols

	// fmt.Printf("Ind: %d, Row: %d, Col: %d\n",ind, row, col)

	return row, col
}

func pointInRectangle(point Vector2, orig Vector2, height int, width int) bool {
	// check x
	px := int(point.X) // annoying to type cast every time
	ox := int(orig.X)  // annoying to type cast every time
	checkX := (px > ox) && (px < (ox + width))

	// check y
	py := int(point.Y) // annoying to type cast every time
	oy := int(orig.Y)  // annoying to type cast every time
	checkY := (py > oy) && (py < (oy + height))

	return checkX && checkY
}
