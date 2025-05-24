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

func DrawDebugPlayerPos(p types.Player){
  positionText := fmt.Sprintf("[X: %.3f | Y: %.3f | Z: %.3f]", p.Pos.X, p.Pos.Y, p.Pos.Z)

  DrawText(positionText, 10, 10, 20, Black)
}
func DrawDebugPlayerFPS(){
  fpsText := fmt.Sprintf("FPS: %d", GetFPS())
  DrawText(fpsText, 10, 30, 20, Black)
}
func DrawDebugActiveBlock(b *types.Block){
  textPos := NewVector2(10, 50)
  if b == nil{
    DrawText("No Active Block", int32(textPos.X), int32(textPos.Y), 20, Black)
    return
  }

  blockText := fmt.Sprintf("Block Type: %d BlockPosition: Chunk Index: [Row: %d, Col: %d] Block Index: [I: %d,J: %d,K: %d]", b.Type,
			b.BlockPosition.ChunkIndex.Row,
			b.BlockPosition.ChunkIndex.Col,
			b.BlockPosition.BlockIndex.I,
			b.BlockPosition.BlockIndex.J,
			b.BlockPosition.BlockIndex.K,
    )
  DrawText(blockText, int32(textPos.X), int32(textPos.Y), 20, Black)
}

func PrintPlayerHand(inv *types.Inventory){
  // item
  if inv.Hand.Item == nil {
    fmt.Println("No Item")
    return
  }
  fmt.Printf("Hand| ID: %d, Count: %d\n", inv.Hand.Item.GetId(), inv.Hand.Count)
}

func DrawDebugBlockPosition(bp *types.BlockPosition, t string, offset int){
	var text string
	if bp == nil {
		text = fmt.Sprintf("%-10s: nil", t)
	}else {
		text = fmt.Sprintf("%-10s: Chunk Index: [Row: %d, Col: %d] Block Index: [I: %d,J: %d,K: %d]",
			t,
			bp.ChunkIndex.Row,
			bp.ChunkIndex.Col,
			bp.BlockIndex.I,
			bp.BlockIndex.J,
			bp.BlockIndex.K,
			)
	}

  DrawText(text, 10, int32(70 + offset), 20, Black)
}
