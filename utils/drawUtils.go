package utils

import (
	"blockProject/constants"
	types "blockProject/types"
	// "fmt"

	. "github.com/gen2brain/raylib-go/raylib"
)

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
  }

  // if b.Focused{
    // DrawCubeWires(centerPoint,
    //   constants.BlockWidth,
    //   constants.BlockHeight,
    //   constants.BlockDepth,
    //   Black)
  // }
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

func DrawChunks(w types.World, p types.Player) {

  idcs := types.GetRenderableChunkIndeces(p, w)

  if len(idcs) == 0{
    return
  }

  // new and improved
  for _, val := range idcs{
    j, i := types.UnboxChunkIndex(val)
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






