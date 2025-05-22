package types

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

// interface
type PlaceableItem interface{
  Item
  GetBlockType() BlockType
  Place(world *World, postion Vector3) bool
}

type BlockItem struct{
  BaseItem
  Type BlockType
}


func (b *BlockItem) GetBlockType() BlockType { return b.Type }

func (b *BlockItem) Place(world *World, chunkInd ChunkIndex, blockInd BlockIndex) bool {
  // validate the position
  //val := world.ValidatePosition()


  // implement later
  return true

}


