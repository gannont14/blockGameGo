package types

import (
	"blockProject/constants"
	"fmt"

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
  BoundBox BoundingBox
}

type BlockIndex struct {
  I int
  J int
  K int
}

func NewBlockIndex(i int, j int, k int) BlockIndex {
  return BlockIndex{
    I: i,
    J: j,
    K: k,
  }
}

func (bi *BlockIndex) UnboxBlockIndex() (int, int, int) {
  return bi.I, bi.J, bi.K
}

func NewBlock(t BlockType, pos Vector3, chunkId int) Block{
  // generate bounding box min and max based on position vector
  bbMin := pos
  bbMax := Vector3Add(bbMin, 
    NewVector3(constants.BlockWidth, 
      constants.BlockHeight, 
      constants.BlockDepth))

  fmt.Printf("Box from [%f, %f, %f] to [%f, %f, %f]\n",
    bbMin.X, bbMin.Y, bbMin.Z, 
    bbMax.X, bbMax.Y, bbMax.Z, 
    )
  // generate actual block struct
  b := Block{
    Type: t,
    WorldPos: pos,
    ChunkId: chunkId,
    Focused: false,
    BoundBox: NewBoundingBox(bbMin, bbMax),
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
