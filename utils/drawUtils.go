package utils

import (
	"blockProject/constants"
	types "blockProject/types"
	// "fmt"

	. "github.com/gen2brain/raylib-go/raylib"
)

func DrawCrosshair() {
  // this is straight from claude, not going to lie
    // Get screen center
    screenWidth := GetScreenWidth()
    screenHeight := GetScreenHeight()
    centerX := screenWidth / 2
    centerY := screenHeight / 2
    
    // Crosshair settings
    crosshairSize := 10
    crosshairThickness := 2
    crosshairColor := Black
    
    // Draw horizontal line
    DrawLineEx(
        Vector2{X: float32(centerX - crosshairSize), Y: float32(centerY)},
        Vector2{X: float32(centerX + crosshairSize), Y: float32(centerY)},
        float32(crosshairThickness),
        crosshairColor,
    )
    
    // Draw vertical line
    DrawLineEx(
        Vector2{X: float32(centerX), Y: float32(centerY - crosshairSize)},
        Vector2{X: float32(centerX), Y: float32(centerY + crosshairSize)},
        float32(crosshairThickness),
        crosshairColor,
    )
}

func drawBlock(b types.Block){
  // find center point of block
  centerPoint := b.CenterPoint()

  c := b.BlockColor()
  if c != White{
    DrawCube(centerPoint,
      constants.BlockWidth,
      constants.BlockHeight,
      constants.BlockDepth,
      b.BlockColor())
    // Draw the bounding box to debug
    // DrawBoundingBox(b.BoundBox, Blue)
  }

  if b.Focused{
    DrawCubeWires(centerPoint,
      constants.BlockWidth,
      constants.BlockHeight,
      constants.BlockDepth,
      Black)
  }
}

func drawChunk(c types.Chunk){
  if(len(c.Blocks)) == 0{
    return 
  }
  count := 0
  for i := range (len(c.Blocks)){
    for j := range(len(c.Blocks[0])){
      for k := range(len(c.Blocks[0][0])){
        b := c.Blocks[i][j][k]
        drawBlock(b)
        count += 1
      }
    }
  }

  // DEBUG, draw borde around chunk
  DrawChunkBorder(c)

  // fmt.Println("Drew Chunks: ", count)
}

func DrawChunks(w types.World, p types.Player, idcs []types.ChunkIndex) {

  // moved out to main function
  // idcs := types.GetRenderableChunkIndeces(p, w)

  if len(idcs) == 0{
    return
  }

  // new and improved
  for _, val := range idcs{
    j, i := val.UnboxChunkIndex()
    c := w.Chunks[i][j]
    drawChunk(c)
  }


  // deprecated
  // for i := range len(w.Chunks){
  //   for j := range len(w.Chunks[0]){
  //     c := w.Chunks[i][j]
  //     if(c.ShouldBeRendered(p)){
  //       drawChunk(c)
  //     }
  //   }
  // }
}






