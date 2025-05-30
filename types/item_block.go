package types

import (
	textures "blockProject/textures"
	"fmt"

	. "github.com/gen2brain/raylib-go/raylib"
)

type BlockHardness int 

const(
  BaseLevel BlockHardness = iota
  WoodLevel
  StoneLevel
  IronLevel
  GoldLevel
  DiamondLevel
)

type BlockItem struct{
  BaseItem
  Type BlockType
	PrefToolType ToolType
  BaseBreakTime float64 // base break time if it is possible to break with tool
  BlockHardness BlockHardness
	Color Color
	AtlasPosition textures.TextureAtlasPosition

	// interaction handler
	Interactable Interactable
}

func (b *BlockItem) Clone() Item {
	clone := *b
	return &clone
}

func (b *BlockItem) Interact(ctx InteractionContext) bool {
  switch ctx.InteractionType{
  // LEFT CLICK - Break the Block
  case LEFT_CLICK:
  // not sure if going to keep the keyPressed for the interaction context,
  //the left click is mostly handled by the breaking manager




  // RIGHT CLICK - Place the item
  case RIGHT_CLICK:
		fmt.Println("Right Click")
		// check if we are hitting an Interactable block
		if ctx.FocusedBlock != nil && ctx.FocusedBlock.BlockItem.Interactable != nil {
			fmt.Println("Interacion")
			// handle the interaction
			return ctx.FocusedBlock.BlockItem.Interactable.OnBlockInteract(ctx)
		}

		// NOW we check if the player has somethign in their hand
		activeItemStack := &ctx.Player.Inventory.Slots[ctx.Player.ActiveItemSlot]

		if(activeItemStack.Item == nil || activeItemStack.Count <= 0){
			return false
		}

		// now check that the thing in their hand is a block
		if blockItem, isBlockItem := activeItemStack.Item.(*BlockItem); isBlockItem {
			// not intuitive at all, need to mess with the return values
			if ctx.Player.PlaceBlock(ctx, blockItem){
				// decrease inventory
				activeItemStack.Count--
				if activeItemStack.Count <= 0 {
					activeItemStack.nillify()
				}
			}
		}

		// if you made it here, the player is holding a tool, and should use that interaction
    return  true
  }

  // implement later
  return false
}


func (b *BlockItem) IsValidTool(item Item) bool{
  breakLevel := ToolLevelBase
  // check if it's even a tool
  tool, isTool := item.(*ToolItem)

  // if it's a tool, it could have a higher break level
  if isTool {
    breakLevel = tool.ToolLevel
  }

  return breakLevel >= ToolLevel(b.BlockHardness)
}
