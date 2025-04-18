package types

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

type Item interface {
  GetID() int
  GetName() string
  GetMaxStackSize() int
}

type BaseItem struct{
  ID int
  Name string
  MaxStackSize int
}

// base item interface methods
func (b *BaseItem) GetID() int { return b.ID }
func (b *BaseItem) GetName() string { return b.Name }
func (b *BaseItem) GetMaxStackSize() int { return b.MaxStackSize }

// helper function to get an unused itemID
var nextItemId int = 0
func GetNewItemID() int {
  ret := nextItemId
  nextItemId++
  return ret
}

// other interfaces for different kinds of blocks
type PlaceableItem interface{
  Item
  GetBlockType() BlockType
  Place(world *World, postion Vector3) bool
}

type ToolItem interface{
  Item
  Use(world *World, targetPos Vector3) bool
  GetDurability() int
  // GetToolType() 
  // set durability
  // ...
}
