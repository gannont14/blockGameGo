package types

import (
	// . "github.com/gen2brain/raylib-go/raylib"
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
}


func (b *BlockItem) Interact(ctx InteractionContext) bool {
  switch ctx.InteractionType{
  // LEFT CLICK - Break the Block
  case LEFT_CLICK:
  // not sure if going to keep the keyPressed for the interaction context,
  //the left click is mostly handled by the breaking manager




  // RIGHT CLICK - Place the item
  case RIGHT_CLICK:
    return ctx.Player.PlaceBlock(ctx, b) // returns true on success or false
  }

  // implement later
  return false
}


func (b *BlockItem) IsValidTool(item Item) bool{
  breakLevel := BreakLevelBase
  // check if it's even a tool
  tool, isTool := item.(*ToolItem)

  // if it's a tool, it could have a higher break level
  if isTool {
    breakLevel = tool.BreakLevel
  }

  return breakLevel >= BreakLevel(b.BlockHardness)
}
