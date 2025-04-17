package types

import (
  . "github.com/gen2brain/raylib-go/raylib"
)

type Player struct{
  Id int;
  Pos Vector3;
  Cam *Camera;
}

func NewPlayer(startPos Vector3, c *Camera) Player{
  p := Player{
    Id: 0,
    Pos: startPos,
    Cam: c,
  }
  return p
}
