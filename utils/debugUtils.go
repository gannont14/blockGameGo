package utils

import (
  "fmt"
  types "blockProject/types"
	. "github.com/gen2brain/raylib-go/raylib"
)

func PrintPlayerPosition(p types.Player){
  fmt.Printf("Player Position: [%f %f %f]\n", 
    p.Pos.X,
    p.Pos.Y,
    p.Pos.Z,)
}

func DrawChunkBorder(c types.Chunk){
  center := c.CenterPoint()
  w := types.ChunkTotalWidth()
  d := types.ChunkTotalDepth()
  h := types.ChunkTotalHeight()
  DrawCubeWires(center, w, h, d, Blue)
}
