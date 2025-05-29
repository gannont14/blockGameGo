package types

import (
	"blockProject/constants"
	// "fmt"
	// "fmt"

	. "github.com/gen2brain/raylib-go/raylib"
)

type BlockType int

const (
  Air BlockType = iota
  RedBlock
  BlueBlock
  GreenBlock
  BrownBlock
  BlackBlock
  //...
)

type Block struct{
  Type BlockType
  WorldPos Vector3
  BlockPosition BlockPosition
  Focused bool
  BoundBox BoundingBox
	BlockItem *BlockItem
	Color Color
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

func NewBlock(t BlockType, pos Vector3, bp BlockPosition, w *World, ) Block{
  // generate bounding box min and max based on position vector
  bbMin := pos
  bbMax := Vector3Add(bbMin, 
    NewVector3(constants.BlockWidth, 
      constants.BlockHeight, 
      constants.BlockDepth))

  // fmt.Printf("Box from [%f, %f, %f] to [%f, %f, %f]\n",
  //   bbMin.X, bbMin.Y, bbMin.Z, 
  //   bbMax.X, bbMax.Y, bbMax.Z, 
  //   )
	// find from registry for the correct color
	color := Blank
	blockItem, exists := w.ItemRegistry.GetBlockByItemType(t)

	if exists == false {
		// fmt.Println("block item not found of type: ", t)
	} else {
		if t != 0 {
			// fmt.Printf("block item found of type: %d\n", t)
			color = blockItem.(*BlockItem).Color
			// fmt.Printf("Block type %d: Found color %v (R:%d G:%d B:%d A:%d)\n", 
				// t, blockItem.Name, color.R, color.G, color.B, color.A)
		}
	}
  // generate actual block struct
  b := Block{
    Type: t,
    WorldPos: pos,
    BlockPosition: bp,
    Focused: false,
    BoundBox: NewBoundingBox(bbMin, bbMax),
		Color: color,
		BlockItem: blockItem.(*BlockItem),
  }
  return b
}

func (old *Block) Replace (new *Block) {
  // replace all attributes
  old.Type          = new.Type
  old.Focused       = new.Focused
  old.Color         = new.Color
	old.BlockItem     = new.BlockItem
  // other attributes that should be replaced in the future
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

func (b *Block) IsReplaceable() bool {
  switch b.Type {
  case Air:
  // ... add more replaceable and fallthroughs
    return true
  default:
    return false
  }
}

func (b *Block) BlockColor() Color{
  // default
  return b.Color
}
