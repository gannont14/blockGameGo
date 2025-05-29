package types

import (
	"blockProject/constants"
	"fmt"

	. "github.com/gen2brain/raylib-go/raylib"
)

type Player struct{
  Id int;
  Pos Vector3;
  World *World;
  Cam *Camera;
  Inventory *Inventory
  CanMoveCamera bool
  ActiveItemSlot int
}

func NewPlayer(startPos Vector3, c *Camera, w *World) Player{
  // find the number of slots the player should have
  numInvSlots := constants.PlayerInventoryCols * constants.PlayerInventoryRows
  activeSlot := constants.PlayerInventoryCols * (constants.PlayerInventoryRows - 1)
  p := Player{
    Id: 0,
    Pos: startPos,
		World: w,
    Cam: c,
    Inventory: NewInventory(numInvSlots, "PlayerInventory"),
    CanMoveCamera: true,
    ActiveItemSlot: activeSlot, // The first item of the players hotbar
  }
  return p
}

func (p *Player) GenerateActiveBlock(activeChunks []*Chunk,
	focusedBlock **Block,
	focusedBlockPosition **BlockPosition,
	potentialBlock **Block,
  potentialBlockPosition **BlockPosition,
  bm *BreakingManager) {
	w := p.World
  // create ray of player in direciton
  playerLookRay := NewRay(
    p.Cam.Position,
    p.scaledPlayerLookDirection(),
  )


  // reset all block focuses
  for _, c := range activeChunks{
    for i := range c.Blocks{
      for j := range c.Blocks[0]{
        for k := range c.Blocks[0][0]{
          c.Blocks[i][j][k].Focused = false
        }
      }
    }
  }

  // could replace with largest float value, whatever that may be
  closestDistance := float32(9999999)
  // due to odd bug with looking toward origin would cause multiple blocks to be selected
	var playerColl RayCollision
  var closeI, closeJ, closeK int
  var blockChunk *Chunk
  blockHit := false
  *focusedBlock = nil
  *focusedBlockPosition = nil
  *potentialBlockPosition = nil
  *potentialBlock = nil
  var coll RayCollision
  // loop through the active chunks to find
  for _, c := range activeChunks{
    // loop through blocks in said chunk 
    for i := range c.Blocks{
      for j := range c.Blocks[0]{
        for k := range c.Blocks[0][0]{

          // skip if transparent, will have to generalize with additional blocks added
          if(c.Blocks[i][j][k].Type == Air){
            continue
          }

          // check if player look array collides
          coll = GetRayCollisionBox(playerLookRay, c.Blocks[i][j][k].BoundBox)

          if coll.Hit{
            if coll.Distance < closestDistance && coll.Distance < constants.PlayerHighlightRange{
              // c.Blocks[i][j][k].Focused = true
              // fmt.Printf("Colliding with type: %d\n", c.Blocks[i][j][k].Type)
              closestDistance = coll.Distance
              closeI, closeJ, closeK = i, j, k
              blockHit = true
              blockChunk = c
							playerColl = coll
            }
          } 
        }
      }
    }
  }

  // NOW we check and highlight the block
	if blockHit {
		blockChunk.Blocks[closeI][closeJ][closeK].Focused = true
		*focusedBlock = &blockChunk.Blocks[closeI][closeJ][closeK]
		// find the face
		b := *focusedBlock

		// set the focusedBlockPosition
		focBlockPosition := w.worldPosToBlockPosition(b.WorldPos)
		*focusedBlockPosition = &focBlockPosition

		orig := b.CenterPoint()
		face := determineHitFace(orig, playerColl.Point)
    potentialBlockPosition := w.GetPotentialBlock(
      *focusedBlock,
      face, **focusedBlock)
    *potentialBlock = w.BlockPositionToBlock(potentialBlockPosition)

    if *potentialBlock != nil {
      (*potentialBlock).BlockPosition = *potentialBlockPosition
    }

    // handle the breaking
    p.handleBreaking(*focusedBlock,
      bm)

		// place the block
		if IsKeyPressed(KeyF) || IsMouseButtonPressed(MouseButtonRight) {
      icx := NewInteractionContext(
        w,
        p,
        *focusedBlock,
        *potentialBlock,
        RIGHT_CLICK,
        )
      activeItemStack := &p.Inventory.Slots[p.ActiveItemSlot]

      if(activeItemStack.Item == nil || activeItemStack.Count <= 0){
        return
      }

			if activeItemStack.Item.Interact(icx) {
        // decrease inventory
        activeItemStack.Count--
        if activeItemStack.Count <= 0 {
          activeItemStack.nillify()
        }
      }
		}
	} else {
    bm.StopBreaking(p)
  }
}

func (p *Player) handleBreaking(
	focusedBlock *Block, 
	breakingManager *BreakingManager) {
	w := p.World

	// Check if left mouse button is held down
	if IsMouseButtonDown(MouseButtonLeft) || IsKeyDown(KeyT) {
		// Are we already breaking this block?
		if breakingManager.IsPlayerBreaking(p) {
			currentTarget := breakingManager.GetBreakingTarget(p)
			if currentTarget != focusedBlock {
				// Player switched to a different block - start breaking the new one
				breakingManager.StopBreaking(p)
				breakingManager.StartBreaking(p, focusedBlock)
			}
			// check to see if the player switch tools
			tool, _ := p.GetActiveItem()
			if tool != breakingManager.GetItemUsed(p) {
				// restart, switched tools
				breakingManager.StopBreaking(p)
				breakingManager.StartBreaking(p, focusedBlock)
			}
		} else {
			// Start breaking this block
			breakingManager.StartBreaking(p, focusedBlock)
		}

		// Update the breaking progress
		result := breakingManager.UpdateBreaking(p, w)

		switch result {
		case BreakingResultComplete:
			fmt.Println("Successfully broke the block!")
		// Block is already handled by the manager

		case BreakingResultInProgress:
			// In progress
			progress := breakingManager.GetProgress(p)
			if constants.DEBUG {
				DrawText(fmt.Sprintf("Breaking: %.1f%%", progress*100), 10, 190, 20, Black)
				
			}

		case BreakingResultNone:
			// SHould never happen?
			fmt.Println("Warning: Breaking result was None")
		}

	} else {
		// Mouse button not held - stop breaking
		breakingManager.StopBreaking(p)
	}
}

func (p *Player) scaledPlayerLookDirection() Vector3{
  // get dir vector
  viewDirection := Vector3Subtract(p.Cam.Target, p.Cam.Position)
  // normalize to 1
  normViewDir := Vector3Normalize(viewDirection)

  return normViewDir
}

func (p *Player) PlaceBlock(ctx InteractionContext, b *BlockItem) bool {
  // validate active item
  // activeItem := ctx.Player.Inventory.Slots[ctx.Player.ActiveItemSlot].Item
  // validate the potential block position
  val := ctx.World.ValidateBlockPlacement(ctx.PotentialBlock.BlockPosition)

  if !val { return false }

	newBlock := NewBlock(b.Type,
		ctx.PotentialBlock.WorldPos,
		ctx.PotentialBlock.BlockPosition,
		ctx.World)

	ctx.PotentialBlock.Replace(&newBlock)

  return true
}

func (p *Player) BreakBlock(info *BreakingInfo, bm *BreakingManager) {
  // copy so that we can do calculations after the block has been broken
  blockCopy := *info.TargetBlock 
  airBlock := NewBlock(Air, info.TargetBlock.WorldPos, info.TargetBlock.BlockPosition, p.World)
  // replace block with air 
  info.TargetBlock.Replace(&airBlock)

  targetBlockItem, _ := bm.ItemRegistry.GetBlockByItemType(blockCopy.Type)
  // handle tool degrading
  activeItem, isTool := p.GetActiveItem()
  itemShouldDrop := targetBlockItem.(*BlockItem).IsValidTool(activeItem)

  // handle block dropping functionality here
  if itemShouldDrop {
    fmt.Println("Dropping Item")
  }

  if isTool {
    // knock down durability
    activeItem.(*ToolItem).DegradeTool()
  }
}

/*
    Gets the tool the user is holding
    Item: What the actual item is 
    bool: Whether or not the item is a tool
*/

func (p *Player) GetActiveItem() (Item, bool) {
  // active slot
  activeSlot := p.Inventory.Slots[p.ActiveItemSlot]

  // Nothing in hand
  if activeSlot.Item == nil || activeSlot.Count <= 0{
    return nil, false
  }

  // something in hand, check if it's a tool
  item := activeSlot.Item

  _, isTool := item.(*ToolItem)

  return item, isTool
}

func (p *Player) UpdatePlayerCamera(cam *Camera){

  camVec := Vector3{
    X: GetMouseDelta().X*constants.PlayerMouseSensitivity,
    Y: GetMouseDelta().Y*constants.PlayerMouseSensitivity,
    Z: 0.0,
  }

  if p.CanMoveCamera == false {
    camVec = Vector3Scale(camVec, 0.0)
  }

  UpdateCameraPro(cam,                     // Absolute mess
    Vector3{
      X: BoolToFloat(IsKeyDown(KeyW) || IsKeyDown(KeyUp))*constants.PlayerMoveSpeed -
      BoolToFloat(IsKeyDown(KeyS) || IsKeyDown(KeyDown))*constants.PlayerMoveSpeed,
      Y: BoolToFloat(IsKeyDown(KeyD) || IsKeyDown(KeyRight))*constants.PlayerMoveSpeed -
      BoolToFloat(IsKeyDown(KeyA) || IsKeyDown(KeyLeft))*constants.PlayerMoveSpeed,
      Z: BoolToFloat(IsKeyDown(KeySpace)) * constants.PlayerMoveSpeed -
      BoolToFloat(IsKeyDown(KeyC) || IsKeyDown(KeyLeftControl)) * constants.PlayerMoveSpeed,
    },
    camVec,
    GetMouseWheelMove()*2.0)    
}

func (p *Player) getHotBarItems() (int, []ItemStack) {
  l := len(p.Inventory.Slots)
  numHBSlots := constants.PlayerInventoryCols
  hotbarItemStacks := p.Inventory.Slots[(l - numHBSlots): l ]

  return numHBSlots, hotbarItemStacks
}

func (p *Player) UpdatePlayerActiveItem(){
  // player pressed 1-9 keys 
  pressed, slot := keyPressToSlot()
  if pressed {
    p.ActiveItemSlot = slot
  }

  numSlots, _ := p.getHotBarItems()

  // cycle keys pressed 
  if IsKeyPressed(KeyX) {
    newSlot := (p.ActiveItemSlot + 1) % numSlots + constants.HotbarOffset
    p.ActiveItemSlot = newSlot
  }

  // cycle keys pressed 
  if IsKeyPressed(KeyZ) {
    // need to add the num slots incase it goes negative
    newSlot := (p.ActiveItemSlot - 1 + numSlots) % numSlots + constants.HotbarOffset
    p.ActiveItemSlot = newSlot
  }
  // fmt.Println("Active player slot: ", p.ActiveItemSlot)
}


func BoolToFloat(b bool) float32{
  if b{
    return 1.0
  }
  return 0.0
}

func keyPressToSlot() (bool, int){

  offset := constants.HotbarOffset

  if IsKeyPressed(KeyOne)   { return true, offset + 0 }
  if IsKeyPressed(KeyTwo)   { return true, offset + 1 }
  if IsKeyPressed(KeyThree) { return true, offset + 2 }
  if IsKeyPressed(KeyFour)  { return true, offset + 3 }
  if IsKeyPressed(KeyFive)  { return true, offset + 4 }
  if IsKeyPressed(KeySix)   { return true, offset + 5 }
  if IsKeyPressed(KeySeven) { return true, offset + 6 }
  if IsKeyPressed(KeyEight) { return true, offset + 7 }
  if IsKeyPressed(KeyNine)  { return true, offset + 8 }

  return false, -1
}


func determineHitFace(blockOrigin Vector3, hitPoint Vector3) int {
  // figure out which axis has highest absolute value 
	if constants.DEBUG {
		DrawText(fmt.Sprintf("hitPoint: [%f, %f, %f]", hitPoint.X, hitPoint.Y, hitPoint.Z), 
			10,
			130,
			20,
			Black,
			)
		DrawText(fmt.Sprintf("blockOrigin: [%f, %f, %f]", blockOrigin.X, blockOrigin.Y, blockOrigin.Z), 
			10,
			150,
			20,
			Black,
			)
	}

  // difference between origin and hitPoint
  diffVector := Vector3Subtract(hitPoint, blockOrigin)

	if constants.DEBUG {
		DrawText(fmt.Sprintf("diffVector: [%f, %f, %f]", diffVector.X, diffVector.Y, diffVector.Z), 
			10,
			170,
			20,
			Black,
			)
	}

	// determine which value of diff vector is exactly even
	if int(diffVector.Z) ==  1 {return constants.BlockHitFacePosZ}
	if int(diffVector.Z) == -1 {return constants.BlockHitFaceNegZ}
	if int(diffVector.X) ==  1 {return constants.BlockHitFacePosX}
	if int(diffVector.X) == -1 {return constants.BlockHitFaceNegX}
	if int(diffVector.Y) ==  1 {return constants.BlockHitFacePosY}
	if int(diffVector.Y) == -1 {return constants.BlockHitFaceNegY}

	// default return
	return constants.BlockHitFacePosY

}



