package types

import (
	constants "blockProject/constants"
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
)

type Chunk struct{
  Blocks [constants.ChunkSizeX][constants.ChunkSizeY][constants.ChunkSizeZ]Block
  Origin Vector3
}

/*
  Generates the indeces within the world struct that should be rendered
*/
type ChunkIndex struct{
  Row int
  Col int
}

func NewChunkIndex(row int, col int) ChunkIndex{
  c := ChunkIndex{
    Row: row,
    Col: col,
  }

  return c
}

func (c *ChunkIndex)UnboxChunkIndex() (int, int) {
  return c.Row, c.Col
}

func GetChunksFromIndeces(idcs []ChunkIndex, w *World) []*Chunk{

  c := make([]*Chunk, 0)

  for _, val := range idcs{
    j, i := val.UnboxChunkIndex()
    c = append(c, &w.Chunks[i][j])
  }

  return c
}

func GenerateTestChunk(orig Vector3) Chunk{
  // count := 0
  var ch Chunk
  ch.Origin = orig
  // generate a test chunk with a layer of red blocks and the rest air
  // x and z will have the same, onyl changes come with the y coord
  for y := range constants.ChunkSizeY{
    for x := range constants.ChunkSizeX{
      for z := range constants.ChunkSizeZ{

        b := NewBlock(
          Air, 
          Vector3Add(ch.Origin, blockOffset(x, y, z)),
          0,
          )

        // hard code to be the floor
        if y == 0{
          b.Type = RedBlock
        }

        // // hard code to test
        // if y == 1{
        //   b.Type = BlueBlock
        // }

        // add to the chunks block list
        ch.Blocks[x][y][z] = b
        // count += 1
      }
    }
  }

  // fmt.Println("Generated ", count)

  return ch
}

func GenerateTestChunks(xChunks int, yChunks int) [][]Chunk{
  c := make([][]Chunk, xChunks)
  for i := range xChunks{
    s := make([]Chunk, 0, yChunks)
    for j := range yChunks{
      zOffset := ChunkTotalWidth() * float32(i)
      xOffset := ChunkTotalDepth() * float32(j)
      o := NewVector3(xOffset, 0.0, zOffset)
      s = append(s, GenerateTestChunk(o))
      fmt.Printf("Chunk [%d, %d] generated\nAt Origin [%f, %f, %f]\n",
        i, j, o.X, o.Y, o.Z)
    }
    c[i] = s
  }
  return c
}

func blockOffset(x int, y int, z int) Vector3{
  v := Vector3{
    X: float32(x) * constants.BlockWidth,
    Y: float32(y) * constants.BlockHeight,
    Z: float32(z) * constants.BlockDepth,
  }
  return v
}

func ChunkTotalWidth() float32 {
  return float32(constants.ChunkSizeX * constants.BlockWidth)
}

func ChunkTotalDepth() float32 {
  return float32(constants.ChunkSizeZ * constants.BlockDepth)
}

func ChunkTotalHeight() float32 {
  return float32(constants.ChunkSizeY * constants.BlockHeight)
}

/*
  Returns a list of ChunkIndexes that should be drawn by the game,
  The function is massive so could probably be shortened
*/

func GetRenderableChunkIndeces(p Player, w World) []ChunkIndex{
  // Find the plyaers coords and what indeces they should be int he world 
  x := int(p.Pos.X / ChunkTotalWidth())
  z := int(p.Pos.Z / ChunkTotalDepth())

  chunkIndeces := make([]ChunkIndex, 0)

  // find teh numChunksX and numChunksY of world

  numChunksX := len(w.Chunks)
  if(numChunksX == 0){ // SHOULD never happen, Should...
    return chunkIndeces
  }

  numChunksY := len(w.Chunks[0])
  if(numChunksY == 0){ //...
    return chunkIndeces
  }

  // get the negative and positive render distances
  for i := (x - constants.RenderDistance); i <= (x + constants.RenderDistance); i++{
    for j := (z - constants.RenderDistance); j <= (z + constants.RenderDistance); j++{
      // potential chunk index
      potentialChunkInd := NewChunkIndex(i, j)

      // Check bounds within the world array, i.e. not < 0 or > numChunksX
      if(potentialChunkInd .validChunkIndex(numChunksX, numChunksY)){
        // add to the list of valid chunk indeces
        chunkIndeces = append(chunkIndeces, potentialChunkInd)
      }
    }
  }

  return chunkIndeces
}

func (c *ChunkIndex) validChunkIndex(numChunksX int, numChunksY int) bool{
  // valid on the X coordinate
  if(c.Col < 0 || c.Col > numChunksX-1){
    return false
  }
  // valid on the Y coordinate
  if(c.Row < 0 || c.Row > numChunksY-1){
    return false
  }

  return true
}

func (c *Chunk) CenterPoint() Vector3 {
  // calculate the center point of a Chunk
  v := NewVector3(
    c.Origin.X + (constants.BlockWidth * constants.ChunkSizeX) / 2.0,
    c.Origin.Y + (constants.BlockHeight * constants.ChunkSizeY) / 2.0,
    c.Origin.Z + (constants.BlockDepth * constants.ChunkSizeZ) / 2.0,
    )
  return v
}

func (c *Chunk) ShouldBeRendered(p Player) bool {
  d := Vector3Distance(p.Pos, c.CenterPoint())

  // convert the render distance to units
  r := float32(constants.RenderDistance) * float32(constants.ChunkSizeX) * float32(constants.BlockWidth)

  if(d > r){
    return false
  }

  return true
}





