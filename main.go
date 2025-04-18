package main

import (
	"blockProject/constants"
	types "blockProject/types"
	. "blockProject/utils"
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
  items "blockProject/items"
)

var camera Camera3D
var player types.Player
var world types.World
var renderedChunkIndeces []types.ChunkIndex
var renderedChunks []*types.Chunk

var focusedBlock *types.Block = nil
var inventoryDisplayed bool = false

func toggleInventoryStatus(){
  if inventoryDisplayed {
    DisableCursor()
    player.CanMoveCamera = true
    inventoryDisplayed = false
    return
  }

  EnableCursor()
  player.CanMoveCamera = false
  inventoryDisplayed = true
  return
}


func RegisterAllItems(r *types.ItemRegistry) {
  // there HAS to be a better way to do this
  r.RegisterItem(items.NewRedBlockItem())
  r.RegisterItem(items.NewBlueBlockItem())
}

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

  // generate the item registry 
  itemRegistry := types.NewItemRegistry()
  RegisterAllItems(itemRegistry)

  world.ItemRegistry = *itemRegistry
}

// all draw functions
func drawGame(){
  // draw the chunks
  DrawChunks(world, player, renderedChunkIndeces)

  //...
}

func drawHud(){
  // draw crosshair
  DrawCrosshair()

  if inventoryDisplayed {
    DrawInventory(*player.Inventory)
    fmt.Println("Displaying inventory")
  }

  // debug menu
  if constants.DEBUG {
    // render coords
    DrawDebugPlayerPos(player)
    // render focused block
    DrawDebugActiveBlock(focusedBlock)
    // render FPS
    DrawDebugPlayerFPS()
  }
}

func updateGame(){
  // upate the players positions
  pos := NewVector3(camera.Position.X, // Player's position is just the camera's position except for different y
    camera.Position.Y - (constants.PlayerHeight / 2),
    camera.Position.Z)
  player.Pos = pos

  // figure out which chunks to render
  renderedChunkIndeces = types.GetRenderableChunkIndeces(player, world)
  renderedChunks = types.GetChunksFromIndeces(renderedChunkIndeces, &world)

  // find the active block that the player is looking at
  player.GenerateActiveBlock(renderedChunks, &focusedBlock)

  //...
}

func main() {
	InitWindow(1600, 900, "raylib [core] example - basic window")
	defer CloseWindow()

	SetTargetFPS(60)
  DisableCursor()

  initGame()

	for !WindowShouldClose() {


    player.UpdatePlayerCamera(&camera)

    if(IsKeyPressed(KeyE)){
      toggleInventoryStatus()
    }


		BeginDrawing()
    ClearBackground(RayWhite)

    updateGame()

    BeginMode3D(camera)

    drawGame()
    // PrintPlayerPosition(player)
    EndMode3D()

    drawHud()

		EndDrawing()
	}
}
