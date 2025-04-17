package types

import (
	"blockProject/constants"
	. "github.com/gen2brain/raylib-go/raylib"
)

type BlockType int

const (
  Air BlockType = iota
  RedBlock
  BlueBlock
  //...
)

type Block struct{
  Type BlockType
  WorldPos Vector3
  ChunkId int
  Focused bool
}

func NewBlock(t BlockType, pos Vector3, chunkId int) Block{
  b := Block{
    Type: t,
    WorldPos: pos,
    ChunkId: chunkId,
    Focused: false,
  }
  return b
}

func (b *Block) CenterPoint() Vector3{
  // calculate the center point of a block
  v := NewVector3(
    b.WorldPos.X + constants.BlockWidth / 2.0,
    b.WorldPos.Y + constants.BlockHeight / 2.0,
    b.WorldPos.Z + constants.BlockDepth / 2.0,
    )

  return v
}

func (b *Block) BlockColor() Color{
  switch b.Type{
  case Air:
    return White
  case RedBlock:
    return Red
  case BlueBlock:
    return Blue
  }

  // default
  return White
}
