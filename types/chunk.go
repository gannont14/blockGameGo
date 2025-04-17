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

func GenerateTestChunk(orig Vector3) Chunk{
  // count := 0
  var ch Chunk
  ch.Origin = orig
  // generate a test chunk with a layer of red blocks and the rest air
  // x and z will have the same, onyl changes come with the y coord
  for y := range constants.ChunkSizeY{
    for x := range constants.ChunkSizeX{
      for z := range constants.ChunkSizeZ{
        b := Block{
          Type: Air,
          WorldPos: Vector3Add(ch.Origin, blockOffset(x, y, z)),
          ChunkId: 0,
        }

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





