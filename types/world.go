package types

import (
	constants "blockProject/constants"

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


func (w *World) PlaceBlock(focusedBlock *Block, face int, newBlock Block) *BlockPosition {
  // Find the chunk and block from the world position
  bp := w.worldPosToBlockPosition(focusedBlock.WorldPos)

  newBp, _ := blockFaceToBlock(bp, face)

  return &newBp
}

func blockFaceToBlock(blockPos BlockPosition, face int)  (BlockPosition, bool) {
  // find what the supposed index would be 
  dX, dY, dZ := 0, 0, 0
  nBlockPos := blockPos

  switch face {
  case constants.BlockHitFaceNegX:
    // subtract from x directoin 
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
  if dX + blockPos.BlockIndex.I > constants.ChunkSizeX { nBlockPos.ChunkIndex.Col += 1 }
  if dX + blockPos.BlockIndex.I < constants.ChunkSizeX { nBlockPos.ChunkIndex.Col -= 1 }
  if dZ + blockPos.BlockIndex.K > constants.ChunkSizeZ { nBlockPos.ChunkIndex.Row += 1 }
  if dZ + blockPos.BlockIndex.K < constants.ChunkSizeZ { nBlockPos.ChunkIndex.Row -= 1 }

  nBlockPos.BlockIndex.I = blockPos.BlockIndex.I + dX
  nBlockPos.BlockIndex.J = blockPos.BlockIndex.J + dY
  nBlockPos.BlockIndex.K = blockPos.BlockIndex.K + dZ

  return nBlockPos, true
}

func (w *World) worldPosToBlockPosition(pos Vector3)  BlockPosition {
  // figure out which chunk the block is in
  cX := int(pos.X) / (constants.ChunkSizeX * constants.BlockWidth)
  cY := int(pos.Z) / (constants.ChunkSizeZ * constants.BlockDepth)
  chunkInd := NewChunkIndex( cX, cY,)

  // now find the block index within that chunk
  chunkOrig := w.Chunks[cX][cY].Origin
  // position of hte block offset by the chunks origin
  offsetPos := Vector3Subtract(pos, chunkOrig)

  bX := int(offsetPos.X / constants.ChunkSizeX)
  bY := int(offsetPos.Y / constants.ChunkSizeY)
  bZ := int(offsetPos.Z / constants.ChunkSizeZ)

  blockInd := NewBlockIndex(bX, bY, bZ)

  blockPos := NewBlockPosition(chunkInd, blockInd)

  return blockPos
}

func (bp *BlockPosition) BlockPosToWorldPos() Vector3 {

}


