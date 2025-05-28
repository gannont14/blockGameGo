package main

import (
	"blockProject/constants"
	gamestate "blockProject/gamestate"
	items "blockProject/items"
	"blockProject/textures"
	types "blockProject/types"
	. "blockProject/utils"
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	targetFPS int32 = 120

	camera               Camera3D
	renderedChunkIndeces []types.ChunkIndex
	renderedChunks       []*types.Chunk

	gs *gamestate.GameState

	focusedBlock           *types.Block         = nil
	focusedBlockPosition   *types.BlockPosition = nil
	potentialBlock         *types.Block         = nil
	potentialBlockPosition *types.BlockPosition = nil
	inventoryDisplayed     bool                 = false
	hotbarDisplayed        bool                 = true
)

func toggleInventoryStatus() {
	player := gs.Player

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

func initGame() {
	// create the texture atlas
	blockAtlas := textures.NewTextureAtlas("textures/atlases/block_atlas2.png",
		5, 1, 16)

	// create the item atlas
	itemAtlas := textures.NewItemAtlas("textures/atlases/item_atlas.png",
		5, 2, 16)

	// init the camera
	camera.Position = NewVector3(0.0, constants.PlayerHeight, 4.0)
	camera.Target = NewVector3(0.0, 2.0, 0.0)
	camera.Up = NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 60.0
	camera.Projection = CameraPerspective
	fmt.Println("Camera generated")

	// generate the item registry
	itemRegistry := types.NewItemRegistry()
	RegisterAllItems(itemRegistry)
	fmt.Println("ItemRegistry generated")

	// Create world
	world := types.World{}
	world.BlockAtlas = &blockAtlas
	world.ItemRegistry = *itemRegistry

	// generate player
	player := types.NewPlayer(camera.Position, &camera, &world)
	fmt.Println("Player generated")

	// generate the worlds test chunk
	world.Chunks = types.GenerateTestChunks(constants.NumChunksX, constants.NumChunksY, &world)
	fmt.Println("Chunks generated")

	// create the block breaking manager
	breakingManager := types.NewBreakingManager(world.ItemRegistry)

	// Initialize the gamestate with all the resources
	gamestate.GameStateInit(&player, &world, breakingManager, &itemAtlas)

	// Set the global gamestate reference
	gs = gamestate.Get()

	AddItemsToPlayerInv()
}

// all draw functions
func drawGame() {
	// draw the chunks
	DrawChunks(renderedChunkIndeces)

	// active block marker
	if potentialBlock != nil {
		DrawSphere(potentialBlock.CenterPoint(),
			0.2,
			Orange)
	}
	//...
}

func drawHud() {
	// Access from gamestate
	player := gs.Player
	breakingManager := gs.BreakingManager

	if inventoryDisplayed {
		// fmt.Println("Displaying inventory")
		DrawInventory(*&player.Inventory)
	}
	if hotbarDisplayed {
		DrawHotbar(*player.Inventory, player.ActiveItemSlot)
	}
	// draw crosshair
	DrawCrosshair()

	// draw item name
	itemNameOrig := Vector2{X: constants.ScreenWidth / 2 , Y: constants.ScreenHeight - 125}
	DrawItemName(itemNameOrig, 30)

	// debug menu
	if constants.DEBUG {
		DrawDebugPlayerPos(*player)
		DrawDebugActiveBlock(focusedBlock)
		DrawDebugBlockPosition(focusedBlockPosition, "Active", 0)
		DrawDebugBlockPosition(potentialBlockPosition, "Potential", 20)
		DrawDebugPlayerFPS()
	}

	if breakingManager.IsPlayerBreaking(player) {
		prog := breakingManager.GetProgress(player)
		DrawProgressBar(NewVector2(800, 500), 10, 40, prog, Gray, Green)
	}
}

func updateGame() {
	player := gs.Player
	world := gs.World
	breakingManager := gs.BreakingManager

	// update the players positions
	pos := NewVector3(camera.Position.X,
		camera.Position.Y-(constants.PlayerHeight/2),
		camera.Position.Z)
	player.Pos = pos

	// figure out which chunks to render
	renderedChunkIndeces = types.GetRenderableChunkIndeces(*player, *world)
	renderedChunks = types.GetChunksFromIndeces(renderedChunkIndeces, world)

	// find the active block that the player is looking at
	player.GenerateActiveBlock(renderedChunks,
		&focusedBlock,
		&focusedBlockPosition,
		&potentialBlock,
		&potentialBlockPosition,
		breakingManager)

	// update what item the player is holding
	player.UpdatePlayerActiveItem()
	//...
}

func main() {
	InitWindow(constants.ScreenWidth, constants.ScreenHeight, "raylib [core] example - basic window")
	defer CloseWindow()

	SetTargetFPS(500)
	DisableCursor()

	initGame()

	for !WindowShouldClose() {
		player := gs.Player // Access player from gamestate

		player.UpdatePlayerCamera(&camera)

		if IsKeyPressed(KeyE) {
			toggleInventoryStatus()
		}

		BeginDrawing()
		ClearBackground(RayWhite)

		updateGame()

		BeginMode3D(camera)
		drawGame()
		EndMode3D()

		drawHud()
		EndDrawing()
	}
}

func AddItemsToPlayerInv() {
	// Access from gamestate
	player := gs.Player
	world := gs.World

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
