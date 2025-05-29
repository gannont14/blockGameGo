package types

import (
	"blockProject/textures"
	"fmt"
)

// . "github.com/gen2brain/raylib-go/raylib"

type Item interface {
  GetId() int 
  GetName() string 
  GetMaxStackSize() int
  GetAtlasPosition() textures.TextureAtlasPosition
	Clone() Item
  Interact(InteractionContext) bool
}

func (b *BaseItem) Interact(ctx InteractionContext) bool {
  return false
}


func (b *BaseItem) GetId() int { return b.Id }
func (b *BaseItem) GetName() string { return b.Name }
func (b *BaseItem) GetMaxStackSize() int { return b.MaxStackSize }
func (b *BaseItem) GetAtlasPosition() textures.TextureAtlasPosition { return b.AtlasPosition }
func (b *BaseItem) Clone() Item {
	clone := *b
	return &clone
}

type BaseItem struct {
  Id int 
  Name string 
  MaxStackSize int 
	AtlasPosition textures.TextureAtlasPosition
}

// helper function to get an unused itemID
var nextItemId int = 0
func GetNewItemID() int {
  ret := nextItemId
  nextItemId++
	fmt.Println("Generated new item with id: ", ret)
  return ret
}

