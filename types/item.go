package types

import (
	// . "github.com/gen2brain/raylib-go/raylib"
)

type Item interface {
  GetId() int 
  GetName() string 
  GetMaxStackSize() int
  Interact(InteractionContext) bool
}

func (b *BaseItem) GetId() int { return b.Id }
func (b *BaseItem) GetName() string { return b.Name }
func (b *BaseItem) GetMaxStackSize() int { return b.MaxStackSize }

type BaseItem struct {
  Id int 
  Name string 
  MaxStackSize int 
}

func (b *BaseItem) Interact(ctx InteractionContext) bool {
  return false
}

// helper function to get an unused itemID
var nextItemId int = 0
func GetNewItemID() int {
  ret := nextItemId
  nextItemId++
  return ret
}

