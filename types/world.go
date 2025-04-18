package types

import (
  // constants "blockProject/constants"
)


// singleton for the world
type World struct{
  Chunks [][]Chunk
  ItemRegistry ItemRegistry
}

