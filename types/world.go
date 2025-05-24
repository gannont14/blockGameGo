package types

import (
	constants "blockProject/constants"
	"fmt"

	. "github.com/gen2brain/raylib-go/raylib"
)

// singleton for the world
type World struct{
  Chunks [][]Chunk
  ItemRegistry ItemRegistry
}

type BlockPosition struct{
  ChunkIndex ChunkIndex
  BlockIndex BlockIndex
}

func NewBlockPosition(ChunkIndex ChunkIndex, BlockIndex BlockIndex) BlockPosition {
  return BlockPosition{
    ChunkIndex: ChunkIndex,
    BlockIndex: BlockIndex,
  }
}

func UnboxBlockPosition(BlockPosition BlockPosition) (ChunkIndex, BlockIndex) {
  return BlockPosition.ChunkIndex, BlockPosition.BlockIndex
}

func validBlockPos(bp BlockPosition) bool {
	// check right nwo if it's negative, change to check if in world bounds
	if bp.ChunkIndex.Row < 0 ||
	   bp.ChunkIndex.Col < 0 ||
	   bp.BlockIndex.I < 0 || 
	   bp.BlockIndex.J < 0 || 
	   bp.BlockIndex.K < 0 { return false }

	return true

}

/*
  Checks to see if the potential block already has something there
  RETURN: boolean if position is valid
*/
func (w *World) ValidateBlockPlacement(bp BlockPosition) bool {
  b := w.BlockPositionToBlock(&bp)

  // check to make sure that the type is a replaceable type,
  // right now, this is just the air block
  if b.IsReplaceable() {
    return true
  }

  return false
}

/*
  Places Block at block position within world
  RETURN: nada
*/
func (w *World) PlaceBlockAtBlockPosition(b *Block, bp *BlockPosition) {
  // set the block at the block positon to the new one
  currentBlock := w.BlockPositionToBlock(bp)

  // replaces all attributes of the pointer
  currentBlock.Replace(b)
}

func (w *World) BlockPositionToBlock(bp *BlockPosition) *Block {
	if !validBlockPos(*bp) { return nil }
	ci := bp.ChunkIndex
	bi := bp.BlockIndex

	return &w.Chunks[ci.Col][ci.Row].Blocks[bi.I][bi.J][bi.K]
}


func (w *World) GetPotentialBlock(focusedBlock *Block, face int, newBlock Block) *BlockPosition {
  // Find the chunk and block from the world position
  bp := w.worldPosToBlockPosition(focusedBlock.WorldPos)

	DrawText(fmt.Sprintf("Face: %d", face), 10, 110, 20, Black)

  newBp, _ := blockFaceToBlock(bp, face)

  return &newBp
}

func blockFaceToBlock(blockPos BlockPosition, face int)  (BlockPosition, bool) {
  // find what the supposed index would be 
  dX, dY, dZ := 0, 0, 0
  nBlockPos := blockPos

  switch face {
  case constants.BlockHitFaceNegX:
    dX -= 1;
  case constants.BlockHitFaceNegY:
    dY -= 1;
  case constants.BlockHitFaceNegZ:
    dZ -= 1;
  case constants.BlockHitFacePosX:
    dX += 1;
  case constants.BlockHitFacePosY:
    dY += 1;
  case constants.BlockHitFacePosZ:
    dZ += 1;
  }

  // determine if that difference is outside of the bounds of the chunk
  // only need to check the X and Z,
  if dX + blockPos.BlockIndex.I >= constants.ChunkSizeX { nBlockPos.ChunkIndex.Row += 1 }
  if dX + blockPos.BlockIndex.I <          0           { nBlockPos.ChunkIndex.Row -= 1 }
  if dZ + blockPos.BlockIndex.K >= constants.ChunkSizeZ { nBlockPos.ChunkIndex.Col += 1 }
  if dZ + blockPos.BlockIndex.K <          0           { nBlockPos.ChunkIndex.Col -= 1 }

	                                                      // ensure it's a positive value
  nBlockPos.BlockIndex.I = ((blockPos.BlockIndex.I + dX) + constants.ChunkSizeX) % constants.ChunkSizeX
  nBlockPos.BlockIndex.J = ((blockPos.BlockIndex.J + dY) + constants.ChunkSizeY) % constants.ChunkSizeY
  nBlockPos.BlockIndex.K = ((blockPos.BlockIndex.K + dZ) + constants.ChunkSizeZ) % constants.ChunkSizeZ

  return nBlockPos, true
}

func (w *World) worldPosToBlockPosition(pos Vector3)  BlockPosition {
  // figure out which chunk the block is in
  cX := pos.X / (constants.ChunkSizeX * constants.BlockWidth)
  cY := pos.Z / (constants.ChunkSizeZ * constants.BlockDepth)
	// as ints for index
	iCX := int(cX)
	iCY := int(cY)
  chunkInd := NewChunkIndex(iCX, iCY)

  // now find the block index within that chunk
  chunkOrig := w.Chunks[iCY][iCX].Origin
  // position of hte block offset by the chunks origin
  offsetPos := Vector3Subtract(pos, chunkOrig)
	// DrawText(fmt.Sprintf("ChunkOrigin: [%f, %f, %f]", chunkOrig.X, chunkOrig.Y, chunkOrig.Z), 
	// 	200,
	// 	200,
	// 	20,
	// 	Black,
	// 	)

  bX := int(offsetPos.X / constants.BlockWidth)  % constants.ChunkSizeX
  bY := int(offsetPos.Y / constants.BlockHeight) % constants.ChunkSizeY
  bZ := int(offsetPos.Z / constants.BlockDepth)  % constants.ChunkSizeZ

  blockInd := NewBlockIndex(bX, bY, bZ)

  blockPos := NewBlockPosition(chunkInd, blockInd)

  return blockPos
}

func (bp *BlockPosition) BlockPosToWorldPos() Vector3 {
	return NewVector3(0, 0, 0)
}
