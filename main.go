package main

import (
	"blockProject/constants"
	types "blockProject/types"
	. "blockProject/utils"
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
  items "blockProject/items"
	"blockProject/textures"
)

var camera Camera3D
var player types.Player
var world types.World
var renderedChunkIndeces []types.ChunkIndex
var renderedChunks []*types.Chunk
var breakingManager types.BreakingManager
var textureAtlas textures.TextureAtlas

var focusedBlock *types.Block = nil
var focusedBlockPosition *types.BlockPosition = nil
var potentialBlock *types.Block = nil
var potentialBlockPosition *types.BlockPosition = nil
var inventoryDisplayed bool = false
var hotbarDisplayed bool = true

func toggleInventoryStatus(){
  if inventoryDisplayed {
    DisableCursor()
    player.CanMoveCamera = true
    inventoryDisplayed = false
    hotbarDisplayed = true
    return
  }

  EnableCursor()
  player.CanMoveCamera = false
  inventoryDisplayed = true
  hotbarDisplayed = false
  return
}


func initGame(){

	// create the texture atlas 
	textureAtlas = textures.NewTextureAtlas("textures/atlases/block_atlas2.png", 
		5, 1, 16)

	// add to world
	world.TextureAtlas = &textureAtlas

  // init the camera
	camera.Position = NewVector3(0.0, constants.PlayerHeight, 4.0) // Camera position
	camera.Target = NewVector3(0.0, 2.0, 0.0)   // Camera looking at point
	camera.Up = NewVector3(0.0, 1.0, 0.0)       // Camera up vector (rotation towards target)
	camera.Fovy = 60.0                             // Camera field-of-view Y
	camera.Projection = CameraPerspective       // Camera projection type
  fmt.Println("Camera generated")

  // generate the item registry 
  itemRegistry := types.NewItemRegistry()
  RegisterAllItems(itemRegistry)

  world.ItemRegistry = *itemRegistry
  fmt.Println("ItemRegistry generated")

  // generate player
  player = types.NewPlayer(camera.Position, &camera, &world)
  fmt.Println("Player generated")

  // generate the worlds test chunk
  world.Chunks = types.GenerateTestChunks(constants.NumChunksX, constants.NumChunksY, &world)
  fmt.Println("Chunks generated")

  // create the block breaking manager
  breakingManager = *types.NewBreakingManager(world.ItemRegistry)

	AddItemsToPlayerInv()
}

// all draw functions
func drawGame(){
  // draw the chunks
  DrawChunks(world, player, renderedChunkIndeces)

	// active block marker
	if potentialBlock != nil{
		DrawSphere(potentialBlock.CenterPoint(),
			0.2,
			Orange)
	}
  //...
}

func drawHud(){
  // draw crosshair
  DrawCrosshair()

  if hotbarDisplayed {
    DrawHotbar(*player.Inventory, player.ActiveItemSlot)
  }

  if inventoryDisplayed {
    DrawInventory(player.Inventory)
    // fmt.Println("Displaying inventory")
  }

  // debug menu
  if constants.DEBUG {
    // render coords
    DrawDebugPlayerPos(player)
    // render focused block
    DrawDebugActiveBlock(focusedBlock)
		// DrawDebugBlockPosition(&focusedBlock.BlockPosition, "TESTING FOCUSED", 80)
		// DrawDebugBlockPosition(&potentialBlock.BlockPosition, "TESTING ACTIVE", 100)
		DrawDebugBlockPosition(focusedBlockPosition, "Active", 0)
		DrawDebugBlockPosition(potentialBlockPosition, "Potential", 20)
		// render FPS
    DrawDebugPlayerFPS()
  }

	if breakingManager.IsPlayerBreaking(&player)  {
		prog := breakingManager.GetProgress(&player)
		DrawProgressBar(NewVector2(800, 500),10, 40, prog, Gray, Green)
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
  player.GenerateActiveBlock(renderedChunks, 
		&focusedBlock,
		&focusedBlockPosition,
		&potentialBlock,
		&potentialBlockPosition,
    &breakingManager)
  // udpate what item the player is holding
  player.UpdatePlayerActiveItem()

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

func AddItemsToPlayerInv(){
  // add random items to the player inventory

  Redblock, _ := world.ItemRegistry.GetItemByName("Red Block")
  player.Inventory.AddItem(Redblock, 63)

  Blueblock, _ := world.ItemRegistry.GetItemByName("Blue Block")
  player.Inventory.AddItem(Blueblock, 63)

  Brownblock, _ := world.ItemRegistry.GetItemByName("Brown Block")
  player.Inventory.AddItem(Brownblock, 63)

  Greenblock, _ := world.ItemRegistry.GetItemByName("Green Block")
  player.Inventory.AddItem(Greenblock, 63)

  Blackblock, _ := world.ItemRegistry.GetItemByName("Black Block")
  player.Inventory.AddItem(Blackblock, 63)


	// tools
  WoodenPickaxe, _ := world.ItemRegistry.GetItemByName("Wooden Pickaxe")
  player.Inventory.AddItem(WoodenPickaxe, 1)

  WoodenShovel, _ := world.ItemRegistry.GetItemByName("Wooden Shovel")
  player.Inventory.AddItem(WoodenShovel, 1)

  WoodenAxe, _ := world.ItemRegistry.GetItemByName("Wooden Axe")
  player.Inventory.AddItem(WoodenAxe, 1)

  DiamondPickaxe, _ := world.ItemRegistry.GetItemByName("Diamond Pickaxe")
  player.Inventory.AddItem(DiamondPickaxe, 1)

  DiamondShovel, _ := world.ItemRegistry.GetItemByName("Diamond Shovel")
  player.Inventory.AddItem(DiamondShovel, 1)

  DiamondAxe, _ := world.ItemRegistry.GetItemByName("Diamond Axe")
  player.Inventory.AddItem(DiamondAxe, 1)

  GoldPickaxe, _ := world.ItemRegistry.GetItemByName("Gold Pickaxe")
  player.Inventory.AddItem(GoldPickaxe, 1)
}

// Ew


func RegisterAllItems(r *types.ItemRegistry) {
  // there HAS to be a better way to do this
  r.RegisterItem(items.NewAirBlockItem())
  r.RegisterItem(items.NewRedBlockItem())
  r.RegisterItem(items.NewBlueBlockItem())
  r.RegisterItem(items.NewGreenBlockItem())
  r.RegisterItem(items.NewBrownBlockItem())
  r.RegisterItem(items.NewBlackBlockItem())
  r.RegisterItem(items.NewWoodenPickaxeItem())
  r.RegisterItem(items.NewStonePickaxeItem())
  r.RegisterItem(items.NewIronPickaxeItem())
  r.RegisterItem(items.NewGoldPickaxeItem())
  r.RegisterItem(items.NewDiamondPickaxeItem())
  r.RegisterItem(items.NewWoodenAxeItem())
  r.RegisterItem(items.NewStoneAxeItem())
  r.RegisterItem(items.NewIronAxeItem())
  r.RegisterItem(items.NewGoldAxeItem())
  r.RegisterItem(items.NewDiamondAxeItem())
  r.RegisterItem(items.NewWoodenShovelItem())
  r.RegisterItem(items.NewStoneShovelItem())
  r.RegisterItem(items.NewIronShovelItem())
  r.RegisterItem(items.NewGoldShovelItem())
  r.RegisterItem(items.NewDiamondShovelItem())
}
