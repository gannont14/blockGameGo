package types

import (
	// . "github.com/gen2brain/raylib-go/raylib"
)



type BlockItem struct{
  BaseItem
  Type BlockType
}


func (b *BlockItem) Interact(ctx InteractionContext) bool {
  switch ctx.InteractionType{
  // LEFT CLICK - Break the Block
  case LEFT_CLICK:




  // RIGHT CLICK - Place the item
  case RIGHT_CLICK:
    // validate active item
    // activeItem := ctx.Player.Inventory.Slots[ctx.Player.ActiveItemSlot].Item
    // validate the potential block position
    val := ctx.World.ValidateBlockPlacement(ctx.PotentialBlock.BlockPosition)

    if !val { return false }


    // now actually place the block and handle inventory
    // ctx.World.PlaceBlockAtBlockPosition(activeItem,
    //   &ctx.PotentialBlock.BlockPosition)

    // block replacement code for now
    ctx.PotentialBlock.Type = b.Type
    ctx.PotentialBlock.Focused = true

    return true
  }



  // implement later
  return false
}


