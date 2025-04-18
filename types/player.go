package types

import (
	"blockProject/constants"
	// "fmt"
	. "github.com/gen2brain/raylib-go/raylib"
)

type Player struct{
  Id int;
  Pos Vector3;
  Cam *Camera;
  Inventory *Inventory
  CanMoveCamera bool
}

func NewPlayer(startPos Vector3, c *Camera) Player{
  // find the number of slots the player should have
  numInvSlots := constants.PlayerInventoryCols * constants.PlayerInventoryRows
  p := Player{
    Id: 0,
    Pos: startPos,
    Cam: c,
    Inventory: NewInventory(numInvSlots, "PlayerInventory"),
    CanMoveCamera: true,
  }
  return p
}

func (p *Player) GenerateActiveBlock(activeChunks []*Chunk, focusedBlock **Block) {

  // create ray of player in direciton
  playerLookRay := NewRay(
    p.Cam.Position,
    p.scaledPlayerLookDirection(),
  )


  // reset all block focuses
  for _, c := range activeChunks{
    for i := range c.Blocks{
      for j := range c.Blocks[0]{
        for k := range c.Blocks[0][0]{
          c.Blocks[i][j][k].Focused = false
        }
      }
    }
  }

  // could replace with largest float value, whatever that may be
  closestDistance := float32(9999999)
  // due to odd bug with looking toward origin would cause multiple blocks to be selected
  var closeI, closeJ, closeK int
  var blockChunk *Chunk
  blockHit := false
  *focusedBlock = nil
  // loop through the active chunks to find
  for _, c := range activeChunks{
    // loop through blocks in said chunk 
    for i := range c.Blocks{
      for j := range c.Blocks[0]{
        for k := range c.Blocks[0][0]{

          // skip if transparent, will have to generalize with additional blocks added
          if(c.Blocks[i][j][k].Type == Air){
            continue
          }

          // check if player look array collides
          coll := GetRayCollisionBox(playerLookRay, c.Blocks[i][j][k].BoundBox)

          if coll.Hit{
            if coll.Distance < closestDistance && coll.Distance < constants.PlayerHighlightRange{
              // c.Blocks[i][j][k].Focused = true
              // fmt.Printf("Colliding with type: %d\n", c.Blocks[i][j][k].Type)
              closestDistance = coll.Distance
              closeI, closeJ, closeK = i, j, k
              blockHit = true
              blockChunk = c
            }
          } 
        }
      }
    }
  }

  // NOW we check and highlight the block
  if blockHit {
    blockChunk.Blocks[closeI][closeJ][closeK].Focused = true
    *focusedBlock = &blockChunk.Blocks[closeI][closeJ][closeK]
  }
}

func (p *Player) scaledPlayerLookDirection() Vector3{
  // get dir vector
  viewDirection := Vector3Subtract(p.Cam.Target, p.Cam.Position)
  // normalize to 1
  normViewDir := Vector3Normalize(viewDirection)

  return normViewDir
}

func (p *Player) UpdatePlayerCamera(cam *Camera){

  camVec := Vector3{
    X: GetMouseDelta().X*constants.PlayerMouseSensitivity,
    Y: GetMouseDelta().Y*constants.PlayerMouseSensitivity,
    Z: 0.0,
  }

  if p.CanMoveCamera == false {
    camVec = Vector3Scale(camVec, 0.0)
  }

  UpdateCameraPro(cam,                     // Absolute mess
    Vector3{
      X: BoolToFloat(IsKeyDown(KeyW) || IsKeyDown(KeyUp))*constants.PlayerMoveSpeed -
      BoolToFloat(IsKeyDown(KeyS) || IsKeyDown(KeyDown))*constants.PlayerMoveSpeed,
      Y: BoolToFloat(IsKeyDown(KeyD) || IsKeyDown(KeyRight))*constants.PlayerMoveSpeed -
      BoolToFloat(IsKeyDown(KeyA) || IsKeyDown(KeyLeft))*constants.PlayerMoveSpeed,
      Z: BoolToFloat(IsKeyDown(KeySpace)) * constants.PlayerMoveSpeed -
      BoolToFloat(IsKeyDown(KeyC) || IsKeyDown(KeyLeftControl)) * constants.PlayerMoveSpeed,
    },
    camVec,
    GetMouseWheelMove()*2.0)    
}


func BoolToFloat(b bool) float32{
  if b{
    return 1.0
  }
  return 0.0
}
