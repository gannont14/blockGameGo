package main

import (
	"blockProject/constants"
	types "blockProject/types"
	. "blockProject/utils"
	"fmt"

	. "github.com/gen2brain/raylib-go/raylib"
)

var camera Camera3D
var player types.Player
var world types.World
var renderedChunkIndeces []types.ChunkIndex
var renderedChunks []*types.Chunk

func initGame(){
  // init the camera
	camera.Position = NewVector3(0.0, constants.PlayerHeight, 4.0) // Camera position
	camera.Target = NewVector3(0.0, 2.0, 0.0)   // Camera looking at point
	camera.Up = NewVector3(0.0, 1.0, 0.0)       // Camera up vector (rotation towards target)
	camera.Fovy = 60.0                             // Camera field-of-view Y
	camera.Projection = CameraPerspective       // Camera projection type
  fmt.Println("Camera generated")

  // generate player
  player = types.NewPlayer(camera.Position, &camera)
  fmt.Println("Player generated")

  // generate the worlds test chunk
  world.Chunks = types.GenerateTestChunks(constants.NumChunksX, constants.NumChunksY)
  fmt.Println("Chunks generated")


}

// all draw functions
func drawGame(){
  // draw the chunks
  DrawChunks(world, player, renderedChunkIndeces)

  //...
}

func updateGame(){
  // upate the players positions
  pos := NewVector3(camera.Position.X, camera.Position.Y - (constants.PlayerHeight/2), camera.Position.Z)
  player.Pos = pos

  // figure out which chunks to render
  renderedChunkIndeces = types.GetRenderableChunkIndeces(player, world)
  renderedChunks = types.GetChunksFromIndeces(renderedChunkIndeces, &world)

  // find the active block that the player is looking at
  player.GenerateActiveBlock(renderedChunks)

  //...
}

func main() {
	InitWindow(1600, 900, "raylib [core] example - basic window")
	defer CloseWindow()

	SetTargetFPS(60)
  DisableCursor()

  initGame()


	for !WindowShouldClose() {


    UpdateCameraPro(&camera,                     // Absolute mess
      Vector3{
        X: BoolToFloat(IsKeyDown(KeyW) || IsKeyDown(KeyUp))*constants.PlayerMoveSpeed -
           BoolToFloat(IsKeyDown(KeyS) || IsKeyDown(KeyDown))*constants.PlayerMoveSpeed,
        Y: BoolToFloat(IsKeyDown(KeyD) || IsKeyDown(KeyRight))*constants.PlayerMoveSpeed -
           BoolToFloat(IsKeyDown(KeyA) || IsKeyDown(KeyLeft))*constants.PlayerMoveSpeed,
        Z: BoolToFloat(IsKeyDown(KeySpace)) * constants.PlayerMoveSpeed -
           BoolToFloat(IsKeyDown(KeyC) || IsKeyDown(KeyLeftControl)) * constants.PlayerMoveSpeed,
      },
      Vector3{
        X: GetMouseDelta().X*constants.PlayerMouseSensitivity,
        Y: GetMouseDelta().Y*constants.PlayerMouseSensitivity,
        Z: 0.0,
      },
      GetMouseWheelMove()*2.0)    


		BeginDrawing()
    ClearBackground(RayWhite)

    updateGame()

    BeginMode3D(camera)

    drawGame()
    // PrintPlayerPosition(player)
    EndMode3D()

		EndDrawing()
	}
}
